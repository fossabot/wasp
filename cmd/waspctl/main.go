package main

import (
	"context"
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vx-labs/wasp/wasp/api"
	"go.uber.org/zap"
)

func seedRand() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
func main() {
	seedRand()
	config := viper.New()
	ctx := context.Background()
	rootCmd := &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			config.BindPFlags(cmd.Flags())
			config.BindPFlags(cmd.PersistentFlags())
		},
	}
	raft := &cobra.Command{
		Use: "raft",
	}
	node := &cobra.Command{
		Use: "node",
	}
	mqtt := &cobra.Command{
		Use: "mqtt",
	}
	mqtt.AddCommand(&cobra.Command{
		Use: "list-sessions",
		Run: func(cmd *cobra.Command, _ []string) {
			conn, l := mustDial(ctx, cmd, config)
			out, err := api.NewMQTTClient(conn).ListSessionMetadatas(ctx, &api.ListSessionMetadatasRequest{})
			if err != nil {
				l.Fatal("failed to list connected sessions", zap.Error(err))
			}
			table := getTable([]string{"ID", "Client ID", "Peer", "Connected Since"}, cmd.OutOrStdout())
			for _, member := range out.GetSessionMetadatasList() {
				table.Append([]string{
					member.GetSessionID(),
					member.GetClientID(),
					fmt.Sprintf("%x", member.GetPeer()),
					time.Since(time.Unix(member.GetConnectedAt(), 0)).String(),
				})
			}
			table.Render()
		},
	})
	raft.AddCommand(&cobra.Command{
		Use: "members",
		Run: func(cmd *cobra.Command, _ []string) {
			conn, l := mustDial(ctx, cmd, config)
			out, err := api.NewRaftClient(conn).GetMembers(ctx, &api.GetMembersRequest{})
			if err != nil {
				l.Fatal("failed to list raft members", zap.Error(err))
			}
			table := getTable([]string{"ID", "Leader", "Address", "Health"}, cmd.OutOrStdout())
			for _, member := range out.GetMembers() {
				healthString := "healthy"
				if !member.IsAlive {
					healthString = "unhealthy"
				}
				table.Append([]string{
					fmt.Sprintf("%x", member.GetID()), fmt.Sprintf("%v", member.GetIsLeader()), member.GetAddress(), healthString,
				})
			}
			table.Render()
		},
	})
	node.AddCommand(&cobra.Command{
		Use: "shutdown",
		Run: func(cmd *cobra.Command, args []string) {
			conn, l := mustDial(ctx, cmd, config)
			_, err := api.NewNodeClient(conn).Shutdown(ctx, &api.ShutdownRequest{})
			if err != nil {
				l.Fatal("failed to shutdown node", zap.Error(err))
			}
			fmt.Println("Shutdown started")
		},
	})
	rootCmd.AddCommand(raft)
	rootCmd.AddCommand(node)
	rootCmd.AddCommand(mqtt)
	rootCmd.PersistentFlags().BoolP("insecure", "k", false, "Disable GRPC client-side TLS validation.")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Increase log verbosity.")
	rootCmd.PersistentFlags().BoolP("use-consul", "c", false, "Use Hashicorp Consul to find Wasp server.")
	rootCmd.PersistentFlags().String("consul-service-name", "wasp", "Consul service name.")
	rootCmd.PersistentFlags().String("consul-service-tag", "rpc", "Consul service tag.")
	rootCmd.PersistentFlags().String("host", "127.0.0.1:1899", "remote GRPC endpoint")
	rootCmd.PersistentFlags().String("rpc-tls-certificate-authority-file", "", "x509 certificate authority used by RPC Client.")
	rootCmd.PersistentFlags().String("rpc-tls-certificate-file", "", "x509 certificate used by RPC Client.")
	rootCmd.PersistentFlags().String("rpc-tls-private-key-file", "", "Private key used by RPC Client.")
	rootCmd.Execute()
}
