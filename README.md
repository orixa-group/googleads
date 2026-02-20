# Google Ads Go SDK

A clean, idiomatic Go client for the Google Ads API.

## Features

- **Service-Oriented**: Organized by Google Ads resources (Campaigns, Budgets, etc.).
- **gRPC Under the Hood**: High-performance communication using the latest Google Ads protobuf definitions.

## Installation

```bash
go get github.com/orixa-group/googleads
```

## Configuration

The client requires a `Config` struct with your Google Ads API credentials:

```go
config := googleads.Config{
    ClientID:        "YOUR_CLIENT_ID",
    ClientSecret:    "YOUR_CLIENT_SECRET",
    RefreshToken:    "YOUR_REFRESH_TOKEN",
    DeveloperToken:  "YOUR_DEVELOPER_TOKEN",
    LoginCustomerID: "YOUR_LOGIN_CUSTOMER_ID", // Optional
}
```

## Usage

### Connecting to the API

```go
ctx := context.Background()
client, err := googleads.Connect(ctx, config)
if err != nil {
    log.Fatal(err)
}
defer client.Close()
```

### Working with Campaigns

#### Listing Campaigns

```go
campaigns, err := client.Campaigns().List(ctx, "CUSTOMER_ID")
if err != nil {
    log.Fatal(err)
}

for _, c := range campaigns {
    fmt.Printf("Campaign: %s (ID: %s)\n", c.GetName(), c.GetId())
}
```

#### Creating a Campaign

First, create a campaign budget, then use its resource name to create the campaign.

```go
// 1. Create a Campaign Budget
campaign := googleads.NewCampaign()
campaign.SetName("Summer Sale 2024")
campaign.SetBudget(1000) // Amount in cents

budget, err := client.CampaignBudgets().Create(ctx, "CUSTOMER_ID", campaign.Budget)
if err != nil {
    log.Fatal(err)
}

// 2. Create the Campaign using the budget's resource name
campaign.SetEnabled(true)
campaign.CampaignBudget = googleads.String(budget.GetResourceName())

createdCampaign, err := client.Campaigns().Create(ctx, "CUSTOMER_ID", campaign)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Created campaign with ID: %s\n", createdCampaign.GetId())
```

### Working with Campaign Budgets

#### Fetching a Budget

```go
budget, err := client.CampaignBudgets().Fetch(ctx, "CUSTOMER_ID", googleads.CampaignBudgetById("456"))
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Budget Amount: %d cents\n", budget.GetAmountCents())
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.