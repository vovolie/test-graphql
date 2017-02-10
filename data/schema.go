package data

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
	"fmt"
)


var nodeDefinitions *relay.NodeDefinitions
var categoryType *graphql.Object
var materialType *graphql.Object

var Schema graphql.Schema

func init() {
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			switch resolvedID.Type {
			case "Category":
				return GetCategoryById(resolvedID.ID), nil
			case "Material":
				return GetMaterialById(resolvedID.ID), nil
			default:
				return nil, errors.New("Unknown node type")
			}
		},

		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *Category:
				return categoryType
			case *Material:
				return materialType
			default:
				return categoryType
			}
		},
	})

	materialConnectionDefinition := relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name: "Material",
		NodeType: materialType,
	})

	categoryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Category", nil),
			"name": &graphql.Field{
				Type: graphql.String,
				Description: "The name of the category.",
			},
			"materials": &graphql.Field{
				Type: materialConnectionDefinition.ConnectionType,
				Args: relay.ConnectionArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := relay.NewConnectionArguments(p.Args)
					materials := []interface{}{}
					if category, ok := p.Source.(*Category); ok {
						typeMaterials := GetMaterialByCategory(category.ID)
						for _, material := range typeMaterials {
							materials = append(materials, material)
						}
					}

					return relay.ConnectionFromArray(materials, args), nil
				},
			},
			"totalCount": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if category, ok := p.Source.(*Category); ok {
						return  len(GetMaterialByCategory(category.ID)), nil
					}
					return 0, nil
				},
			},
		},

		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},

	})

	materialType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Material",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Material", nil),
			"categoryInfo": &graphql.Field{
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

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"category": &graphql.Field{
				Type: categoryType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetCategory(), nil
				},
			},
			"material": &graphql.Field{
				Type: materialType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetMaterial(), nil
				},
			},
			"node": nodeDefinitions.NodeField,
		},
	})

	var err error

	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	if err != nil {
		panic(err)
	}

}