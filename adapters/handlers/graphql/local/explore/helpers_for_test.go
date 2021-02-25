//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package explore

import (
	"context"
	"net/http"

	"github.com/graphql-go/graphql"
	testhelper "github.com/semi-technologies/weaviate/adapters/handlers/graphql/test/helper"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/modulecapabilities"
	"github.com/semi-technologies/weaviate/entities/moduletools"
	"github.com/semi-technologies/weaviate/entities/search"
	modcontextionaryneartext "github.com/semi-technologies/weaviate/modules/text2vec-contextionary/neartext"
	"github.com/semi-technologies/weaviate/usecases/traverser"
)

type mockRequestsLog struct{}

func (m *mockRequestsLog) Register(first string, second string) {
}

type mockResolver struct {
	testhelper.MockResolver
}

type fakeModulesProvider struct{}

func (p *fakeModulesProvider) ExploreArguments(schema *models.Schema) map[string]*graphql.ArgumentConfig {
	txt2vec := &mockText2vecContextionaryModule{}
	for _, c := range schema.Classes {
		if c.Vectorizer == txt2vec.Name() {
			return txt2vec.ExploreArguments()
		}
	}
	return map[string]*graphql.ArgumentConfig{}
}

func (p *fakeModulesProvider) ExtractSearchParams(arguments map[string]interface{}) map[string]interface{} {
	exractedParams := map[string]interface{}{}
	if param, ok := arguments["nearText"]; ok {
		exractedParams["nearText"] = extractNearTextParam(param.(map[string]interface{}))
	}
	return exractedParams
}

func extractNearTextParam(param map[string]interface{}) interface{} {
	txt2vec := &mockText2vecContextionaryModule{}
	return txt2vec.ExtractFunctions()["nearText"](param)
}

func getFakeModulesProvider() ModulesProvider {
	return &fakeModulesProvider{}
}

func newMockResolver() *mockResolver {
	field := Build(testhelper.SimpleSchema.Objects, getFakeModulesProvider())
	mocker := &mockResolver{}
	mockLog := &mockRequestsLog{}
	mocker.RootFieldName = "Explore"
	mocker.RootField = field
	mocker.RootObject = map[string]interface{}{
		"Resolver":    Resolver(mocker),
		"RequestsLog": mockLog,
	}
	return mocker
}

func newMockResolverNoModules() *mockResolver {
	field := Build(testhelper.SimpleSchema.Objects, nil)
	mocker := &mockResolver{}
	mockLog := &mockRequestsLog{}
	mocker.RootFieldName = "Explore"
	mocker.RootField = field
	mocker.RootObject = map[string]interface{}{
		"Resolver":    Resolver(mocker),
		"RequestsLog": mockLog,
	}
	return mocker
}

func newMockResolverEmptySchema() *mockResolver {
	field := Build(&models.Schema{}, getFakeModulesProvider())
	mocker := &mockResolver{}
	mockLog := &mockRequestsLog{}
	mocker.RootFieldName = "Explore"
	mocker.RootField = field
	mocker.RootObject = map[string]interface{}{
		"Resolver":    Resolver(mocker),
		"RequestsLog": mockLog,
	}
	return mocker
}

func (m *mockResolver) Explore(ctx context.Context,
	principal *models.Principal, params traverser.ExploreParams) ([]search.Result, error) {
	args := m.Called(params)
	return args.Get(0).([]search.Result), args.Error(1)
}

type mockText2vecContextionaryModule struct{}

func (m *mockText2vecContextionaryModule) Name() string {
	return "text2vec-contextionary"
}

func (m *mockText2vecContextionaryModule) Init(params moduletools.ModuleInitParams) error {
	return nil
}

func (m *mockText2vecContextionaryModule) RootHandler() http.Handler {
	return nil
}

func (m *mockText2vecContextionaryModule) GetArguments(classname string) map[string]*graphql.ArgumentConfig {
	return modcontextionaryneartext.New().GetArguments(classname)
}

func (m *mockText2vecContextionaryModule) ExploreArguments() map[string]*graphql.ArgumentConfig {
	return modcontextionaryneartext.New().ExploreArguments()
}

func (m *mockText2vecContextionaryModule) ExtractFunctions() map[string]modulecapabilities.ExtractFn {
	return modcontextionaryneartext.New().ExtractFunctions()
}

func (m *mockText2vecContextionaryModule) VectorSearches() map[string]modulecapabilities.VectorForParams {
	return map[string]modulecapabilities.VectorForParams{}
}
