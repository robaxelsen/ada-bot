// TODO: Remaining huzzah.go causes runtime error

package botReactions

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"github.com/bwmarrin/discordgo"
	"github.com/adayoung/ada-bot/settings"
)

type artie struct {
    Trigger string
}

type Artefact struct {
    ID           int    `json:"id"`
    Title        string `json:"title"`
    Description  string `json:"description"`
    Cost         string `json:"cost"`
    Tags         string `json:"tags"`
}

func (p Artefact) toString() string {
    return toJson(p)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}

func (p *artie) Help() string {
	return "Lookup <artefact> and get details."
}

func (p *artie) HelpDetail() string {
	return p.Help()
}

func (p *artie) Reaction(m *discordgo.Message, a *discordgo.Member, mType string) Reaction {
    w := &tabwriter.Writer{}
    buf := &bytes.Buffer{}
    keyword := m.Content[len(settings.Settings.Discord.BotPrefix)+len(p.Trigger)+1:]
    artefacts := getArtefacts()
    for _, p := range artefacts {
        if strings.Contains(p.Tags, keyword) {
            w.Init(buf, 0, 4, 0, ' ', 0)
            fmt.Fprintf(w, "```\n")
            fmt.Fprintf(w, "Title: \t%s\n", p.Title)
            fmt.Fprintf(w, "Description: \t%s\n", p.Description)
            fmt.Fprintf(w, "Cost: \t%s\n", p.Cost)
            fmt.Fprintf(w, "\n```")
            w.Flush()
	}
    }

    out := buf.String()
    return Reaction{Text: out}
}

func getArtefacts() []Artefact {
    var c []Artefact
    raw, err := ioutil.ReadFile("./utils/data/artefacts.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    json.Unmarshal(raw, &c)
    return c
}

func init() {
    _artie := &artie{
        Trigger: "artefact",
    }
    addReaction(_artie.Trigger, "CREATE", _artie)
}
