package main

import (
	"github.com/anaregdesign/custom-skill-builder/service"
	"github.com/anaregdesign/custom-skill-builder/skill"
	"github.com/anaregdesign/custom-skill-openai/openai"
	"os"
)

func main() {

	resourceName := os.Getenv("AOAI_RESOURCE_NAME")
	deploymentName := os.Getenv("AOAI_DEPLOYMENT_NAME")
	apiVersion := os.Getenv("AOAI_API_VERSION")
	apiKey := os.Getenv("AOAI_API_KEY")

	client := openai.New(resourceName, deploymentName, apiVersion, apiKey)

	embeddingSkill := skill.NewSkill(client.Embed)
	summarizeJaSkill := skill.NewSkill(client.SummarizeJA)
	summarizeEnSkill := skill.NewSkill(client.SummarizeEN)
	summarizeZhSkill := skill.NewSkill(client.SummarizeZH)
	summarizeKoSkill := skill.NewSkill(client.SummarizeKO)
	summarizeFrSkill := skill.NewSkill(client.SummarizeFR)
	summarizeItSkill := skill.NewSkill(client.SummarizeIT)
	summarizeEsSkill := skill.NewSkill(client.SummarizeES)
	summarizePtSkill := skill.NewSkill(client.SummarizePT)
	summarizeRuSkill := skill.NewSkill(client.SummarizeRU)
	summarizeTrSkill := skill.NewSkill(client.SummarizeTR)
	summarizeViSkill := skill.NewSkill(client.SummarizeVI)
	summarizeIdSkill := skill.NewSkill(client.SummarizeID)
	summarizeThSkill := skill.NewSkill(client.SummarizeTH)

	book := skill.NewBook()
	book.Register("embed", embeddingSkill.Flatten())
	book.Register("summarize_ja", summarizeJaSkill.Flatten())
	book.Register("summarize_en", summarizeEnSkill.Flatten())
	book.Register("summarize_zh", summarizeZhSkill.Flatten())
	book.Register("summarize_ko", summarizeKoSkill.Flatten())
	book.Register("summarize_fr", summarizeFrSkill.Flatten())
	book.Register("summarize_it", summarizeItSkill.Flatten())
	book.Register("summarize_es", summarizeEsSkill.Flatten())
	book.Register("summarize_pt", summarizePtSkill.Flatten())
	book.Register("summarize_ru", summarizeRuSkill.Flatten())
	book.Register("summarize_tr", summarizeTrSkill.Flatten())
	book.Register("summarize_vi", summarizeViSkill.Flatten())
	book.Register("summarize_id", summarizeIdSkill.Flatten())
	book.Register("summarize_th", summarizeThSkill.Flatten())

	svc := service.NewCustomSkillService(book)
	svc.Run()

}
