# 🧠 CSum: CLI AI Toolchain 

![Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?logo=go&logoColor=white)
![AI](https://img.shields.io/badge/Powered%20by-Gemini%20AI-4285F4?logo=google&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-CLI%20Tool-orange)
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
go build -o csum
```

🏁 This will create a binary named `csum` in your directory.

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
./csum --input article.txt
```

### 💾 Save Output to File:
```bash
./csum --input article.txt --output summary.txt
```

### 🎛️ Adjust Summary Length:
```bash
./csum --input report.txt --length short
./csum --input research.txt --length medium
./csum --input research.txt --length long
```

### 🆘 View Help:
```bash
./csum --help
```

---

## ⚙️ CLI Flags

| 🏷️ Flag | 💡 Description | 🧩 Example |
|----------|----------------|-------------|
| `--input` | Path to the input text file | `--input notes.txt` |
| `--output` | Path to save the summarized text | `--output summary.txt` |
| `--length` | Summary detail level (`short`, `medium`, `long`) | `--length medium` |

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
🧩 [GoFlag](https://pkg.go.dev/flag) — For argument parsing.  

---

## 🧩 Example Command Flow

```bash
# 🏗️ Build the project
go build -o csum

# ⚙️ Summarize and save output
./csum --input blog.txt --output summary.txt --length medium
```

---

## 💬 Feedback

💡 Have ideas or suggestions? Feel free to open an issue or start a discussion!  
📫 You can also connect via [LinkedIn](https://linkedin.com/in/uddeshya-singh-9b3049236/) or open a PR directly.

---

> 🧩 *Made with ❤️ using Go and Gemini AI API*

