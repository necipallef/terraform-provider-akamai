package property

// This file contains functions removed from resource_akamai_property.go that are still referenced elsewhere

import (
	"context"
	"fmt"
	"strconv"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v7/pkg/papi"
	"github.com/akamai/terraform-provider-akamai/v5/pkg/meta"
	"github.com/akamai/terraform-provider-akamai/v5/pkg/tools"
)

func getGroup(ctx context.Context, meta meta.Meta, groupID string) (*papi.Group, error) {
	logger := meta.Log("PAPI", "getGroup")
	client := Client(meta)
	logger.Debugf("Fetching groups")
	res, err := client.GetGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrFetchingGroups, err.Error())
	}
	groupID = tools.AddPrefix(groupID, "grp_")

	var group *papi.Group
	var groupFound bool
	for _, g := range res.Groups.Items {
		if g.GroupID == groupID {
			group = g
			groupFound = true
			break
		}
	}
	if !groupFound {
		return nil, fmt.Errorf("%w: %s", ErrGroupNotFound, groupID)
	}
	logger.Debugf("Group found: %s", group.GroupID)
	return group, nil
}

func getContract(ctx context.Context, meta meta.Meta, contractID string) (*papi.Contract, error) {
	logger := meta.Log("PAPI", "getContract")
	client := Client(meta)
	logger.Debugf("Fetching contract")
	res, err := client.GetContracts(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrFetchingContracts, err.Error())
	}
	var contract *papi.Contract
	var contractFound bool
	for _, c := range res.Contracts.Items {
		if c.ContractID == contractID {
			contract = c
			contractFound = true
			break
		}
	}
	if !contractFound {
		return nil, fmt.Errorf("%w: %s", ErrContractNotFound, contractID)
	}

	logger.Debugf("Contract found: %s", contract.ContractID)
	return contract, nil
}

func getProduct(ctx context.Context, meta meta.Meta, productID, contractID string) (*papi.ProductItem, error) {
	logger := meta.Log("PAPI", "getProduct")
	client := Client(meta)
	if contractID == "" {
		return nil, ErrNoContractProvided
	}
	logger.Debugf("Fetching product")
	res, err := client.GetProducts(ctx, papi.GetProductsRequest{ContractID: contractID})
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrProductFetch, err.Error())
	}
	var productFound bool
	var product papi.ProductItem
	for _, p := range res.Products.Items {
		if p.ProductID == productID {
			product = p
			productFound = true
			break
		}
	}
	if !productFound {
		return nil, fmt.Errorf("%w: %s", ErrProductNotFound, productID)
	}

	logger.Debugf("Product found: %s", product.ProductID)
	return &product, nil
}

func convertString(v string) interface{} {
	if f1, err := strconv.ParseFloat(v, 64); err == nil {
		return f1
	}
	// FIXME: execution will never reach this as every int representation will be captured by ParseFloat() above
	// this should either be moved above ParseFloat block or removed
	if f2, err := strconv.ParseInt(v, 10, 64); err == nil {
		return f2
	}
	if f3, err := strconv.ParseBool(v); err == nil {
		return f3
	}
	return v
}

func findProperty(ctx context.Context, name string, meta meta.Meta) (*papi.Property, error) {
	client := Client(meta)
	results, err := client.SearchProperties(ctx, papi.SearchRequest{Key: papi.SearchKeyPropertyName, Value: name})
	if err != nil {
		return nil, err
	}
	if len(results.Versions.Items) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrPropertyNotFound, name)
	}

	property, err := client.GetProperty(ctx, papi.GetPropertyRequest{
		ContractID: results.Versions.Items[0].ContractID,
		GroupID:    results.Versions.Items[0].GroupID,
		PropertyID: results.Versions.Items[0].PropertyID,
	})
	if err != nil {
		return nil, err
	}
	if len(property.Properties.Items) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrPropertyNotFound, name)
	}
	return property.Properties.Items[0], nil
}
