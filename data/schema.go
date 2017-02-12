package data

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)


var nodeDefinitions *relay.NodeDefinitions
var materialsConnection *relay.GraphQLConnectionDefinitions


var categoryType *graphql.Object
var materialType *graphql.Object

var Schema graphql.Schema


func init() {

	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			if resolvedID.Type == "Material" {
				return GetMaterial(resolvedID.ID), nil
			}
			if resolvedID.Type == "Category" {
				return GetCategory(resolvedID.ID), nil
			}

			return nil, nil
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *Material:
				return materialType
			case *Category:
				return categoryType
			default:
				return categoryType
			}
		},
	})

	materialType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Material",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Material", nil),
			"categoryinfo": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"cover": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	materialsConnection = relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name: "Material",
		NodeType: materialType,
	})

	categoryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Category", nil),
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"materils": &graphql.Field{
				Type: materialsConnection.ConnectionType,
				Args: relay.NewConnectionArgs(graphql.FieldConfigArgument{
					"categoryId": &graphql.ArgumentConfig{
						Type: graphql.String,
						DefaultValue: "any",
					},
				}),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Source.(*Category).ID
					args := relay.NewConnectionArguments(p.Args)
					materials := MaterialsToSliceInterface(GetMaterials(id))
					return relay.ConnectionFromArray(materials, args), nil
				},
			},
			"totalCount": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return len(GetMaterials("any")), nil
				},
			},
			"current": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Source.(*Category).ID
					return len(GetMaterials(id)), nil
				},
			},
		},
	})

	rootType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Root",
		Fields: graphql.Fields{
			"viewer": &graphql.Field{
				Type: categoryType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetViewer(), nil
				},
			},
			"node": nodeDefinitions.NodeField,
		},
	})

	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootType,
	})
	if err != nil {
		panic(err)
	}

}