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
		Name:     "Material",
		NodeType: materialType,
	})

	categoryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Category", nil),
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"materials": &graphql.Field{
				Type: materialsConnection.ConnectionType,
				Args: relay.ConnectionArgs,
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
					return GetViewer("1"), nil
				},
			},
			"node": nodeDefinitions.NodeField,
		},
	})

	addMaterialMutation := relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "AddMaterial",
		InputFields: graphql.InputObjectConfigFieldMap{
			"categoryInfo": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"cover": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"url": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		OutputFields: graphql.Fields{
			"materialEdge": &graphql.Field{
				Type: materialsConnection.EdgeType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					payload, _ := p.Source.(map[string]interface{})
					materialId, _ := payload["materialId"].(string)
					material := GetMaterial(materialId)
					return relay.EdgeType{
						Node:   material,
						Cursor: relay.CursorForObjectInConnection(MaterialsToSliceInterface(GetMaterials(material.CategoryInfo)), material),
					}, nil
				},
			},
			"viewer": &graphql.Field{
				Type: categoryType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetViewer(p.Source.(*Category).ID), nil
				},
			},
		},
		MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
			categoryInfo, _ := inputMap["categoryInfo"].(string)
			name, _ := inputMap["name"].(string)
			cover, _ := inputMap["cover"].(string)
			url, _ := inputMap["url"].(string)
			materialId := AddMaterial(categoryInfo, name, cover, url)
			return map[string]interface{}{
				"materialId": materialId,
			}, nil
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addMaterial": addMaterialMutation,
		},
	})

	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootType,
		Mutation: mutationType,
	})
	if err != nil {
		panic(err)
	}

}
