# 🦉 The Ravensfield Collection

[The Ravensfield Collection](https://ravensfield.art) is a weird, wonderful, and fully AI-generated museum 🔮  
Every object and every story in this collection is the result of machine imagination.

---

## 💡 Why I Built This Project

🤖 I’m fascinated by the boundaries between AI and creativity.  
💻 I’m self-taught, so everything is great practice.  
👻 And frankly, I love making weird stuff.

---

## ⚙️ Tech Stack

- **Backend**: Go. Why Go? I just _love_ the language and wanted to build an API with it.
- **Frontend**: Ghost CMS with a custom theme. I didn’t want to get bogged down in frontend work, and Ghost streamlined the process perfectly.
- **AI**: [Claude API](https://www.anthropic.com/api) for text, [Leonardo.Ai](https://leonardo.ai/api/) for images
- **Image Hosting**: [Cloudinary](https://cloudinary.com/)

---

## 🧪 How It Works

1. A prompt is generated using a Mad Libs-style function, combining concepts from curated artwork-related lists.
2. The prompt is passed to Leonardo.Ai, which returns a unique image. That image is uploaded to Cloudinary.
3. I send the image link to [Claude’s vision endpoint](https://docs.anthropic.com/en/docs/build-with-claude/vision), with detailed instructions for how to interpret and describe it.
   - _Claude even generates fake quotes from imaginary experts. How cool is that?_
4. The generated text then goes through **two custom Claude message requests**:
   - A "style matcher" that compares the article to three of my own and adjusts the tone to better match my authorial voice.
   - A virtual editor that polishes and formats the content.
5. The final draft is submitted to Ghost CMS in "draft" mode.  
   This is the **human-in-the-loop** moment — the publisher can review, tweak, and schedule content for release.
6. All of this is automated via a **GitHub Actions CI/CD pipeline** triggered by a scheduled cron job.
   - You can also manually hit endpoints or pause the pipeline at any time.

---

## 🚀 Deployment

The Go backend is deployed to a single [Fly.io](https://fly.io/) container.

It’s super cost-efficient: the server sleeps until woken by a cron job or manual request.

---

## 🔮 What’s Next?

There’s still tons of room to grow. Here’s what I’d like to explore next:

- Improve prompts to generate more unique, varied results
- Reduce repetition and robotic tone in Claude’s writing
- Move off Ghost and build a custom frontend
- Try alternative LLMs for experimentation
- Train a model specifically tuned for this kind of creative generation

---

## 🙏 Acknowledgements

Huge thanks to the maintainers of these open-source projects:

- [gosec](https://github.com/securego/gosec)
- [staticcheck](https://github.com/dominikh/go-tools)
- [chi](https://github.com/go-chi/chi)
- [godotenv](https://github.com/joho/godotenv)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [blackfriday](https://github.com/russross/blackfriday)
- [go-webp](https://github.com/kolesa-team/go-webp)

---

🪄 Enjoy wandering the halls of the Ravensfield museum!
