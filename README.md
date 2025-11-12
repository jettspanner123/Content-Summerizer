# 🧠 ContentSummarizerCLI

A **lightweight and efficient command-line tool** that automatically summarizes lengthy text files into concise, easy-to-read summaries — powered by **Go** and **Gemini AI API**.  
Designed for developers, writers, and researchers who want to quickly reduce content size **without losing key insights**.

---

## ✨ Features

- ⚡ **Fast & Lightweight:** Built using Go for maximum performance.  
- 🧩 **AI-Powered Summaries:** Uses Google’s **Gemini AI** API to intelligently condense content.  
- 📄 **File Support:** Read input from a text file and export the summary to another file.  
- 🛠️ **Customizable:** Supports adjustable summary length and tone.  
- 🧰 **Simple CLI Interface:** Easy to use with minimal commands.

---

## 🚀 Installation

Make sure you have **Go 1.21+** installed.

```bash
git clone https://github.com/yourusername/ContentSummarizerCLI.git
cd ContentSummarizerCLI
go build -o summarize
```

This will create a binary called `summarize` in your project directory.

---

## 🔑 Configuration

Before running, set your **Gemini API key** as an environment variable:

### Linux / macOS:
```bash
export GEMINI_API_KEY="your_api_key_here"
```

### Windows (PowerShell):
```powershell
setx GEMINI_API_KEY "your_api_key_here"
```

---

## 🧩 Usage

### Basic Example:
Summarize a text file and print the result to the terminal:
```bash
./summarize -input article.txt
```

### Save Output to File:
```bash
./summarize -input article.txt -output summary.txt
```

### Control Summary Length:
```bash
./summarize -input paper.txt -length short
./summarize -input paper.txt -length detailed
```

### View Help:
```bash
./summarize -help
```

---

## ⚙️ CLI Flags

| Flag | Description | Example |
|------|--------------|----------|
| `-input` | Path to the input text file | `-input notes.txt` |
| `-output` | Path to save the summarized text | `-output summary.txt` |
| `-length` | Summary detail level (`short`, `medium`, `detailed`) | `-length short` |
| `-apiKey` | (Optional) Provide Gemini API key inline | `-apiKey <your_key>` |

---

## 🧠 How It Works

1. Reads content from the specified file.  
2. Sends it to **Gemini AI** for summarization using a tuned prompt.  
3. Receives and formats the summary.  
4. Outputs it to the terminal or an output file.  

---

## 📦 Example Output

**Input:**
> Artificial Intelligence (AI) has emerged as one of the most transformative forces...

**Output:**
> AI is revolutionizing industries through automation and intelligent decision-making, offering both opportunities and challenges in ethics and workforce adaptation.

---

## 🧑‍💻 Built With

- [Go](https://go.dev/) – Fast and efficient CLI foundation  
- [Gemini AI API](https://ai.google.dev/) – For natural language summarization  
- [Cobra](https://github.com/spf13/cobra) or [flag](https://pkg.go.dev/flag) – Command-line argument handling  

---

## 🤝 Contributing

Contributions are welcome!  
1. Fork the repo  
2. Create a new branch  
3. Make your changes  
4. Submit a pull request  

---

## 📜 License

This project is licensed under the **MIT License**.  
See [LICENSE](LICENSE) for details.

---

## 🧩 Example Command Flow

```bash
# Build the project
go build -o summarize

# Summarize and save output
./summarize -input blog.txt -output summary.txt -length medium
```

