<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>LightVault – Secure .env Manager</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      max-width: 800px;
      margin: 2rem auto;
      padding: 0 1rem;
      color: #333;
    }
    h1 {
      color: #2c3e50;
    }
    h2 {
      color: #34495e;
      border-bottom: 1px solid #ddd;
      padding-bottom: 4px;
    }
    pre {
      background: #f4f4f4;
      padding: 1rem;
      overflow-x: auto;
    }
    code {
      font-family: monospace;
      background: #f0f0f0;
      padding: 2px 4px;
    }
    a {
      color: #2980b9;
    }
  </style>
</head>
<body>
  <h1>🚀 LightVault</h1>
  <p><strong>LightVault</strong> is an open-source developer tool for securely managing your <code>.env</code> secrets locally. Inspired by Infisical, it helps you safely store, version, and share environment variables for development projects.</p>

  <h2>✨ Features</h2>
  <ul>
    <li>Secure local secret management</li>
    <li>Simple CLI – easy to set up</li>
    <li>Version control friendly</li>
    <li>Open-source and community-driven</li>
  </ul>

  <h2>⚡ Installation</h2>
  <p>Clone the repo and build it:</p>
  <pre><code>git clone https://github.com/namamishra786/lightvault.git
cd lightvault
go build -o lightvault</code></pre>

  <h2>🚀 Usage</h2>
  <p>Example commands:</p>
  <pre><code># Initialize LightVault
./lightvault init

# Add a new secret
./lightvault add DB_PASSWORD

# List all secrets
./lightvault list

# Export to .env
./lightvault export</code></pre>

  <h2>📂 Project Structure</h2>
  <ul>
    <li><code>main.go</code> – CLI entry point</li>
    <li><code>cmd/</code> – CLI commands</li>
    <li><code>vault/</code> – Secret handling logic</li>
  </ul>

  <h2>🛠️ Contributing</h2>
  <p>Contributions are welcome! Feel free to fork the repo, open issues, and create pull requests. 🙌</p>

  <h2>🔗 Links</h2>
  <ul>
    <li><a href="https://github.com/namamishra786/lightvault" target="_blank">GitHub Repository</a></li>
    <li><a href="https://namanmishra786.vercel.app/" target="_blank">My Portfolio</a></li>
  </ul>

  <h2>👨‍💻 Author</h2>
  <p>Built by <a href="https://www.linkedin.com/in/naman-mishra-5386a9227/" target="_blank">Naman Mishra</a>. <br/>
  Feel free to connect on LinkedIn or contribute to the project!</p>

  <h2>📜 License</h2>
  <p>MIT License – free to use, modify, and distribute.</p>
</body>
</html>
