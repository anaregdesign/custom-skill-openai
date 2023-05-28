# custom-skill-openai
Web server for [Custom Web API Skill](https://learn.microsoft.com/en-us/azure/search/cognitive-search-custom-skill-web-api) for [Azure Cognitive Search](https://learn.microsoft.com/en-us/azure/search/search-what-is-azure-search) that uses [OpenAI](https://openai.com/) to embedding and summarizing text.

## Container
[latestImage](https://github.com/anaregdesign/custom-skill-openai/pkgs/container/custom-skill-openai)

```bash
docker pull ghcr.io/anaregdesign/custom-skill-openai:latest
```

```bash
docker run -p 8080:8080 ghcr.io/anaregdesign/custom-skill-openai:latest
```

### Environments
* `AOAI_RESOURCE_NAME`
* `AOAI_DEPLOYMENT_NAME`
* `AOAI_API_VERSION`
* `AOAI_API_KEY`

### Port
This API runs on port `8080`.

### Endoints
* Embedding
  * `/skills/embed`
* Summarization
  * `/skills/summarize_ja`
  * `/skills/summarize_en`
  * `/skills/summarize_zh`
  * `/skills/summarize_ko`
  * `/skills/summarize_fr`
  * `/skills/summarize_it`
  * `/skills/summarize_es`
  * `/skills/summarize_pt`
  * `/skills/summarize_ru`
  * `/skills/summarize_tr`
  * `/skills/summarize_vi`
  * `/skills/summarize_id`
  * `/skills/summarize_th`

## Example
### Request

```http request
POST http://localhost:8080/skills/summarize_en
Content-Type: application/json

{
    "values": [
        {
            "recordId": "a1",
            "data": {
                "input": "I am not unmindful that some of you have come here out of great trials and tribulations. Some of you have come fresh from narrow jail cells. And some of you have come from areas where your quest -- quest for freedom left you battered by the storms of persecution and staggered by the winds of police brutality. You have been the veterans of creative suffering. Continue to work with the faith that unearned suffering is redemptive. Go back to Mississippi, go back to Alabama, go back to South Carolina, go back to Georgia, go back to Louisiana, go back to the slums and ghettos of our northern cities, knowing that somehow this situation can and will be changed. Let us not wallow in the valley of despair, I say to you today, my friends. And so even though we face the difficulties of today and tomorrow, I still have a dream. It is a dream deeply rooted in the American dream. I have a dream that one day this nation will rise up and live out the true meaning of its creed: \"We hold these truths to be self-evident, that all men are created equal.\" I have a dream that one day on the red hills of Georgia, the sons of former slaves and the sons of former slave owners will be able to sit down together at the table of brotherhood. I have a dream that one day even the state of Mississippi, a state sweltering with the heat of injustice, sweltering with the heat of oppression, will be transformed into an oasis of freedom and justice. I have a dream that my four little children will one day live in a nation where they will not be judged by the color of their skin but by the content of their character. I have a dream today! I have a dream that one day, down in Alabama, with its vicious racists, with its governor having his lips dripping with the words of \"interposition\" and \"nullification\" -- one day right there in Alabama little black boys and black girls will be able to join hands with little white boys and white girls as sisters and brothers. I have a dream today!"
            }
        }
    ]
}
```

### Response

```http request
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 28 May 2023 03:35:15 GMT
Content-Length: 796
Connection: close

{
  "values": [
    {
      "recordId": "a1",
      "data": {
        "output": "The speaker acknowledges the struggles and hardships faced by those in attendance, including imprisonment and persecution. Despite these challenges, he encourages them to continue their pursuit of freedom and to have faith that their suffering will ultimately lead to redemption. He urges them to return to their communities with the belief that change is possible. The speaker then shares his dream of a future where all people are treated equally, regardless of their race, and where former slaves and slave owners can sit together as equals. He envisions a future where his children will not be judged by the color of their skin, but by their character. He dreams of a day when black and white children can join hands as brothers and sisters."
      }
    }
  ]
}```

