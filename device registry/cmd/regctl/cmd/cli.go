package cmd

import (
	"context"
	"github.com/piusalfred/registry"
	"github.com/piusalfred/registry/api"
	"github.com/piusalfred/registry/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	ErrWTF = errors.New("wtf is this: it should never happen")
)

type ReqType int

const (
	Add ReqType = iota
	Get
	Delete
	List
	Update
)

type CLI interface {
	UsersCmd(ctx context.Context, reqType ReqType) func(cmd *cobra.Command, args []string)
	NodesCmd(ctx context.Context, reqType ReqType) func(cmd *cobra.Command, args []string)
	RegionsCmd(ctx context.Context, reqType ReqType) func(cmd *cobra.Command, args []string)
}

type list struct {
	endpoints api.Endpoints
}

func (l list) UsersCmd(ctx context.Context, reqType ReqType) func(cmd *cobra.Command, args []string) {
	switch reqType {
	case List:
		return func(cmd *cobra.Command, args []string) {
			users, err := l.endpoints.ListUser(ctx)
			if err != nil {
				logError(err)
			}

			logJSON(users)
		}

	case Get:
		return func(cmd *cobra.Command, args []string) {
			id, err := cmd.Flags().GetString("id")

			if err != nil || id == "" {
				logUsage(cmd.Short)
				return
			}

			user, err := l.endpoints.GetUser(ctx, id)
			if err != nil {
				logError(err)
				return
			}

			logJSON(user)
		}

	case Add:
		return func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("name")
			email, err := cmd.Flags().GetString("email")
			password, err := cmd.Flags().GetString("password")
			region, err := cmd.Flags().GetString("region")

			if err != nil || name == "" || email == "" ||
				password == "" || region == "" {
				logUsage(cmd.Short)
				return
			}

			user := registry.User{
				Name:     name,
				Email:    email,
				Password: password,
				Region:   region,
			}

			err = l.endpoints.AddUser(ctx, user)

			if err != nil {
				logError(err)
				return
			}

			logCreated("ok. new User successfully created")
		}

	case Delete:
		return func(cmd *cobra.Command, args []string) {
			id, err := cmd.Flags().GetString("id")

			if err != nil || id == "" {
				logUsage(cmd.Short)
				return
			}

			err = l.endpoints.DeleteUser(ctx, id)
			if err != nil {
				logError(err)
				return
			}

			logOK()
		}

	case Update:
		return func(cmd *cobra.Command, args []string) {
			id, err := cmd.Flags().GetString("id")
			group, err := cmd.Flags().GetInt("group")
			region, err := cmd.Flags().GetString("region")

			if err != nil || (region == "" && (group <= 0 || group > 3)) {
				logUsage(cmd.Example)
				return
			}

			user := registry.User{
				ID:     id,
				Group:  group,
				Region: region,
			}

			up, err := l.endpoints.UpdateUser(context.Background(), id, user)

			if err != nil {
				logError(err)
				return
			}

			logCreated("new user created")
			logJSON(up)
		}

	default:
		return func(cmd *cobra.Command, args []string) {
			logUsage(cmd.Short)
		}
	}

}

func (l list) NodesCmd(ctx context.Context, reqType ReqType) func(cmd *cobra.Command, args []string) {

	switch reqType {
	case List:
		return func(cmd *cobra.Command, args []string) {
			nodes, err := l.endpoints.ListNodes(ctx)
			if err != nil {
				logError(err)
			}

			logJSON(nodes)
		}

	case Delete:
		return func(cmd *cobra.Command, args []string) {
			id, err := cmd.Flags().GetString("id")

			if err != nil || id == "" {
				logUsage(cmd.Short)
				return
			}

			err = l.endpoints.DeleteNode(ctx, id)
			if err != nil {
				logError(err)

			} else {
				logOK()
			}

		}

	case Get:
		return func(cmd *cobra.Command, args []string) {
			id, err := cmd.Flags().GetString("id")

			if err != nil || id == "" {
				logUsage(cmd.Short)
				return
			}

			node, err := l.endpoints.GetNode(ctx, id)
			if err != nil {
				logError(err)
				return
			}

			logJSON(node)
		}

	case Add:
		return func(cmd *cobra.Command, args []string) {
			//adr,name,region,lat,long,master,type
			addr, err := cmd.Flags().GetString("adr")
			name, err := cmd.Flags().GetString("name")
			region, err := cmd.Flags().GetString("region")
			latd, err := cmd.Flags().GetString("lat")
			long, err := cmd.Flags().GetString("long")
			master, err := cmd.Flags().GetString("master")
			typ, err := cmd.Flags().GetInt("type")

			if err != nil || name == "" || addr == "" ||
				latd == "" || region == "" || long == "" {
				logUsage(cmd.Short)
				return
			}

			node := registry.Node{
				Addr:   addr,
				Name:   name,
				Type:   typ,
				Region: region,
				Latd:   latd,
				Long:   long,
				Master: master,
			}

			err = l.endpoints.AddNode(ctx, node)

			if err != nil {
				logError(err)
				return
			}

			logCreated("ok. new User successfully created")
		}

	default:
		return func(cmd *cobra.Command, args []string) {
			logUsage(cmd.Short)
		}
	}

}

func (l list) RegionsCmd(ctx context.Context, reqType ReqType) func(cmd *cobra.Command, args []string) {
	switch reqType {
	case List:
		return func(cmd *cobra.Command, args []string) {
			regions, err := l.endpoints.ListRegions(ctx)
			if err != nil {
				logError(err)
			}

			logJSON(regions)
		}

	case Add:
		return func(cmd *cobra.Command, args []string) {
			name, err := cmd.Flags().GetString("name")
			id, err := cmd.Flags().GetString("id")
			description, err := cmd.Flags().GetString("desc")

			if err != nil || name == "" || id == "" || description == "" {
				logUsage(cmd.Short)
			}

			region := registry.Region{
				ID:   id,
				Name: name,
				Desc: description,
			}

			err = l.endpoints.AddRegion(context.Background(), region)

			if err != nil {
				logError(err)
			}

			logCreated("new region added")
		}

	default:
		return func(cmd *cobra.Command, args []string) {
			logError(ErrWTF)
		}
	}
}

func cli(addr, port string) (CLI, error) {
	endpoints, err := api.MakeClientEndpoints(addr + port)

	if err != nil {
		return nil, err
	}
	return list{endpoints: endpoints}, nil
}
