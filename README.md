# 🧠 ContentSummarizerCLI

![Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?logo=go&logoColor=white)
![AI](https://img.shields.io/badge/Powered%20by-Gemini%20AI-4285F4?logo=google&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Platform](https://img.shields.io/badge/Platform-CLI%20Tool-orange)
![Contributions](https://img.shields.io/badge/Contributions-Welcome-blue)
![Build](https://img.shields.io/badge/Build-Passing-brightgreen)
![Status](https://img.shields.io/badge/Status-Active-success)
![Go Version](https://img.shields.io/badge/Go-1.21+-blue)

> 🚀 A **lightweight and efficient command-line tool** that automatically summarizes lengthy text files into concise, easy-to-read summaries — powered by **Go** and **Gemini AI API**.  
> 🧩 Designed for developers, writers, and researchers who want to quickly reduce content size **without losing key insights**.

---

## ✨ Features

🌟 **Fast & Lightweight:** Built using Go for maximum performance.  
🧠 **AI-Powered Summaries:** Uses Google’s Gemini AI API for intelligent content reduction.  
📁 **File Input/Output:** Read content from one file and export summarized results to another.  
🎚️ **Adjustable Summary Length:** Control depth — short, medium, or detailed.  
🧰 **Simple CLI Interface:** Straightforward and easy to use.  
🔒 **Secure:** Keeps your API key private and configurable.  

---

## ⚙️ Installation

Make sure you have **Go 1.21+** installed.

```bash
git clone https://github.com/yourusername/ContentSummarizerCLI.git
cd ContentSummarizerCLI
go build -o summarize
```

🏁 This will create a binary named `summarize` in your directory.

---

## 🔑 Configuration

Before running, set your **Gemini API key** as an environment variable:

### 🐧 Linux / 🍎 macOS:
```bash
export GEMINI_API_KEY="your_api_key_here"
```

### 🪟 Windows (PowerShell):
```powershell
setx GEMINI_API_KEY "your_api_key_here"
```

---

## 💻 Usage Examples

### 🧾 Basic Command:
```bash
./summarize -input article.txt
```

### 💾 Save Output to File:
```bash
./summarize -input article.txt -output summary.txt
```

### 🎛️ Adjust Summary Length:
```bash
./summarize -input report.txt -length short
./summarize -input research.txt -length detailed
```

### 🆘 View Help:
```bash
./summarize -help
```

---

## ⚙️ CLI Flags

| 🏷️ Flag | 💡 Description | 🧩 Example |
|----------|----------------|-------------|
| `-input` | Path to the input text file | `-input notes.txt` |
| `-output` | Path to save the summarized text | `-output summary.txt` |
| `-length` | Summary detail level (`short`, `medium`, `detailed`) | `-length medium` |
| `-apiKey` | (Optional) Provide Gemini API key inline | `-apiKey <your_key>` |

---

## 🧠 How It Works

1. 📥 Reads your text file.  
2. 🔗 Sends the content to **Gemini AI** for summarization.  
3. 🧾 Receives and formats the output.  
4. 💾 Displays or saves the summary based on user flags.  

---

## 📦 Example Output

**📝 Input:**
> Artificial Intelligence (AI) has emerged as one of the most transformative forces...

**⚡ Output:**
> AI is revolutionizing industries through automation and intelligent decision-making, offering both opportunities and challenges in ethics and workforce adaptation.

---

## 🧑‍💻 Built With

🦋 [Go](https://go.dev/) — For a fast and efficient CLI foundation.  
🤖 [Gemini AI API](https://ai.google.dev/) — For natural language summarization.  
🧩 [Cobra](https://github.com/spf13/cobra) or [flag](https://pkg.go.dev/flag) — For argument parsing.  

---

## 🤝 Contributing

🙌 Contributions are always welcome!

1. 🍴 Fork the repo  
2. 🌿 Create a new branch  
3. 💻 Make your changes  
4. 🔁 Submit a pull request  

---

## 📜 License

🪪 This project is licensed under the **MIT License**.  
See [LICENSE](LICENSE) for details.

---

## 🧩 Example Command Flow

```bash
# 🏗️ Build the project
go build -o summarize

# ⚙️ Summarize and save output
./summarize -input blog.txt -output summary.txt -length medium
```

---

## 💬 Feedback

💡 Have ideas or suggestions? Feel free to open an issue or start a discussion!  
📫 You can also connect via [LinkedIn](https://linkedin.com) or open a PR directly.

---

> 🧩 *Made with ❤️ using Go and Gemini AI API*

