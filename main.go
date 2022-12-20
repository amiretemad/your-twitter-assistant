package main

/*
func main() {

	os.Exit(0)
	config := oauth1.NewConfig("", "")
	token := oauth1.NewToken("", "")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	update, _, err := client.Statuses.Update("Just a test :)", nil)
	if err != nil {
		panic(err)
		return
	}

	print(update)

	/*	var a []string
			strings := append(a, "EtemadAmir")

			lookup, h, err := client.Users.Lookup(&twitter.UserLookupParams{
				ScreenName: strings,
			})

			if err != nil {
				return
			}

			all, err := io.ReadAll(h.Body)
			if err != nil {
				return
			}
			fmt.Println(all)
			fmt.Println(lookup)
		get, h, err := client.DirectMessages.Get(nil)
		if err != nil {
			return
		}
		print(h, get)

		items, h, err := client.Friends.List(nil)

		if err != nil {
			panic(err)
			return
		}

		print(items)
		body, error := io.ReadAll(h.Body)
		if error != nil {
			fmt.Println(error)
		}

		fmt.Println((body))
	}
*/
