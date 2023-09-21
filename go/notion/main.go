package main

import (
	"context"
	"os"

	"github.com/jomei/notionapi"
)

func main() {
	token := notionapi.Token(os.Getenv("NOTION_TOKEN"))
	databaseId := notionapi.DatabaseID(os.Getenv("NOTION_DATABASE_ID"))
	client := notionapi.NewClient(token)
	resp, err := client.Database.Query(context.Background(), databaseId, &notionapi.DatabaseQueryRequest{
		Sorts: []notionapi.SortObject{{
			Property:  "Created",
			Direction: notionapi.SortOrderDESC,
		}},
	},
	)
	if err != nil {
		println(err.Error())
		return
	}

	for i := 0; i < len(resp.Results); i++ {
		println(resp.Results[i].CreatedTime.Date())
	}
	println(len(resp.Results))
	println(resp.Results[0].CreatedTime.Date())
	println(resp.Results[len(resp.Results)-1].CreatedTime.Date())
	println(resp.NextCursor.String())

	if resp.HasMore {
		res, _ := client.Database.Query(context.Background(), databaseId, &notionapi.DatabaseQueryRequest{
			StartCursor: resp.NextCursor,
		})
		println(res.Results[len(res.Results)-1].CreatedTime.Date())

		if res.HasMore {
			re, _ := client.Database.Query(context.Background(), databaseId, &notionapi.DatabaseQueryRequest{
				StartCursor: res.NextCursor,
			})
			println(re.Results[len(re.Results)-1].CreatedTime.Date())
			println(len(re.Results))
		}
	}
}
