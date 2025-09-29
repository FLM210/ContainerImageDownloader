package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	LISTENPORT := os.Getenv("LISTENPORT")
	http.HandleFunc("/", handleForm)
	http.HandleFunc("/download", handleDownload)
	if LISTENPORT != "" {
		log.Fatal(http.ListenAndServe(":"+LISTENPORT, nil))
	} else {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Container Image Downloader</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    <style>
        :root {
            --primary-color: #4a90e2;
            --secondary-color: #357ac0;
            --success-color: #155724;
            --error-color: #721c24;
            --light-bg: #f5f7fa;
            --card-bg: #ffffff;
            --text-color: #333333;
            --text-secondary: #555555;
            --border-color: #dddddd;
            --shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            --radius: 10px;
        }

        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        }

        body {
            background-color: var(--light-bg);
            background-image: radial-gradient(circle at 25% 25%, rgba(74, 144, 226, 0.05) 0%, rgba(255, 255, 255, 0) 50%),
                              radial-gradient(circle at 75% 75%, rgba(74, 144, 226, 0.05) 0%, rgba(255, 255, 255, 0) 50%);
            background-attachment: fixed;
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: center;
            min-height: 100vh;
            padding: 20px;
        }

        .container {
            background-color: var(--card-bg);
            border-radius: var(--radius);
            box-shadow: var(--shadow);
            padding: 40px;
            width: 100%;
            max-width: 600px;
            transition: transform 0.3s ease;
        }

        .container:hover {
            transform: translateY(-5px);
        }

        .logo-container {
            text-align: center;
            margin-bottom: 30px;
        }

        .logo {
            font-size: 48px;
            color: var(--primary-color);
            margin-bottom: 10px;
        }

        h1 {
            color: var(--text-color);
            margin-bottom: 20px;
            text-align: center;
            font-weight: 600;
        }

        .subtitle {
            text-align: center;
            color: var(--text-secondary);
            margin-bottom: 30px;
        }

        .form-group {
            margin-bottom: 25px;
        }

        label {
            display: block;
            margin-bottom: 8px;
            font-weight: 600;
            color: var(--text-secondary);
        }

        input[type="text"], select {
            width: 100%;
            padding: 14px;
            border: 1px solid var(--border-color);
            border-radius: 6px;
            font-size: 16px;
            transition: border-color 0.3s, box-shadow 0.3s;
        }

        input[type="text"] {
            overflow-x: auto;
        }

        input[type="text"]:focus, select:focus {
            border-color: var(--primary-color);
            outline: none;
            box-shadow: 0 0 0 3px rgba(74, 144, 226, 0.2);
        }

        .checkbox-group {
            display: flex;
            align-items: center;
            margin-top: 10px;
        }

        .checkbox-group input {
            margin-right: 10px;
            width: auto;
        }

        button {
            background-color: var(--primary-color);
            color: white;
            border: none;
            border-radius: 6px;
            padding: 14px 20px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            width: 100%;
            transition: background-color 0.3s, transform 0.2s;
            display: flex;
            align-items: center;
            justify-content: center;
        }

        button:hover {
            background-color: var(--secondary-color);
            transform: translateY(-2px);
        }

        button:active {
            transform: translateY(0);
        }

        .message {
            margin-top: 20px;
            padding: 15px;
            border-radius: 6px;
            text-align: center;
            transition: all 0.3s ease;
        }

        .success {
            background-color: #d4edda;
            color: var(--success-color);
            border: 1px solid #c3e6cb;
        }

        .error {
            background-color: #f8d7da;
            color: var(--error-color);
            border: 1px solid #f5c6cb;
        }

        .info {
            background-color: #e3f2fd;
            color: #0d47a1;
            border: 1px solid #bbdefb;
        }

        .loading {
            display: inline-block;
            width: 20px;
            height: 20px;
            border: 3px solid rgba(255, 255, 255, 0.3);
            border-radius: 50%;
            border-top-color: white;
            animation: spin 1s ease-in-out infinite;
            margin-right: 10px;
        }

        @keyframes spin {
            to { transform: rotate(360deg); }
        }

        .footer {
            margin-top: 30px;
            color: var(--text-secondary);
            font-size: 14px;
            text-align: center;
        }

        .history-section {
            margin-top: 40px;
            padding-top: 20px;
            border-top: 1px solid var(--border-color);
        }

        .history-title {
            display: flex;
            align-items: center;
            color: var(--text-color);
            margin-bottom: 15px;
            font-weight: 600;
        }

        .history-title i {
            margin-right: 10px;
        }

        .history-list {
            max-height: 150px;
            overflow-y: auto;
        }

        .history-item {
            padding: 12px;
            border-radius: 6px;
            background-color: rgba(74, 144, 226, 0.05);
            margin-bottom: 10px;
            cursor: pointer;
            transition: background-color 0.3s;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .history-item:hover {
            background-color: rgba(74, 144, 226, 0.1);
        }

        .history-item .copy-btn {
            background: none;
            border: none;
            color: var(--primary-color);
            cursor: pointer;
            padding: 5px;
            width: auto;
        }

        .history-item .copy-btn:hover {
            color: var(--secondary-color);
        }

        .placeholder {
            text-align: center;
            color: var(--text-secondary);
            padding: 20px;
        }

        @media (max-width: 600px) {
            .container {
                padding: 25px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="logo-container">
            <div class="logo"><i class="fas fa-boxes"></i></div>
            <h1>Container Image Downloader</h1>
        </div>
        <p class="subtitle">Easily download container images from any registry</p>

        <form id="downloadForm" action="/" method="post">
            <div class="form-group">
                <label for="image">Image Name</label>
                <input type="text" id="image" name="image" required oninput="this.title = this.value;">
            </div>

            <div class="form-group">
                <label for="format">Output Format</label>
                <select id="format" name="format">
                    <option value="docker-archive">Docker Archive</option>
                    <option value="oci-archive">OCI Archive</option>
                </select>
            </div>

            <button type="submit" id="submitBtn">
                <i class="fas fa-download"></i> Download Image
            </button>
        </form>

        <div id="message" class="message" style="display: none;"></div>

    </div>

    <div class="footer">
        <p>Container Image Downloader © 2023</p>
    </div>

    <script>
        document.addEventListener('DOMContentLoaded', function() {
            const downloadForm = document.getElementById('downloadForm');
            const imageInput = document.getElementById('image');
            const messageDiv = document.getElementById('message');
            const submitBtn = document.getElementById('submitBtn');

            // Load download history from localStorage

            downloadForm.addEventListener('submit', function(e) {
                // Basic validation
                if (!imageInput.value.trim()) {
                    e.preventDefault();
                    showMessage('Please enter an image name', 'error');
                    return;
                }

                // Show loading state
                submitBtn.disabled = true;
                submitBtn.innerHTML = '<span class="loading"></span>Downloading...';
                messageDiv.style.display = 'none';
            });

            function showMessage(text, type) {
                messageDiv.textContent = text;
                messageDiv.className = 'message ' + type;
                messageDiv.style.display = 'block';

                // Auto-hide after 5 seconds
                setTimeout(() => {
                    messageDiv.style.display = 'none';
                }, 5000);
            }



            function truncateString(str, maxLength) {
                if (str.length <= maxLength) {
                    return str;
                }
                return str.substring(0, maxLength) + '...';
            }

            // Check for URL parameters to show messages and trigger download
            const urlParams = new URLSearchParams(window.location.search);
            if (urlParams.has('success')) {
                const filename = urlParams.get('filename');
                // showMessage('', 'success');
                // Reset submit button state
                submitBtn.disabled = false;
                submitBtn.innerHTML = '<i class="fas fa-download"></i> Download Image';
                
                // Trigger file download if filename is present
                if (filename) {
                    setTimeout(() => {
                        const downloadUrl = '/download?filename=' + encodeURIComponent(filename);
                        window.location.href = downloadUrl;
                    }, 1000);
                }
            } else if (urlParams.has('error')) {
                showMessage('Failed to download image: ' + urlParams.get('error'), 'error');
                // Reset submit button state
                submitBtn.disabled = false;
                submitBtn.innerHTML = '<i class="fas fa-download"></i> Download Image';
            }
        });
    </script>
</body>
</html>`)

	} else if r.Method == "POST" {
		arr := strings.Split(r.FormValue("image"), "/")
		filename := strings.Replace(arr[len(arr)-1], ":", "_", 1) + ".tar"
		if r.FormValue("format") == "docker-archive" {
			filename = "oci-archive_" + filename
		}

		savetag := strings.Join(arr[1:], "/")

		image := r.FormValue("image")

		_, err := os.Stat(filename)
		if err != nil {
			Command("skopeo copy docker://" + image + " " + r.FormValue("format") + ":" + filename + ":" + savetag + " --src-tls-verify=false --dest-compress-format zstd --dest-compress-level 5")
		}
		// 重定向到带有成功参数和文件名的URL
		http.Redirect(w, r, "/?success=true&filename="+filename, http.StatusSeeOther)
	}
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		showMessage(w, "Filename not specified", "error")
		return
	}

	_, err := os.Stat(filename)
	if err != nil {
		showMessage(w, "File not found: "+filename, "error")
		return
	}

	// 设置下载头信息
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/tar")
	http.ServeFile(w, r, filename)
}

func showMessage(w http.ResponseWriter, text string, typeStr string) {
	// 这个函数用于在服务端渲染错误信息
	// 实际项目中可能需要更复杂的模板渲染
	fmt.Fprintf(w, `<html><body><div class="message %s">%s</div></body></html>`, typeStr, text)
}

func Command(cmd string) error {
	//c := exec.Command("cmd", "/C", cmd) 	// windows
	fmt.Println(cmd)
	c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, stdErr := c.StdoutPipe()
	if stdErr != nil {
		return stdErr
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(2)
	// 处理标准输出
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Println(readString)
		}
	}()
	// 处理标准错误
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stderr)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Fprintln(os.Stderr, readString)
		}
	}()
	// 启动命令
	err = c.Start()
	if err != nil {
		return err
	}
	// 等待输出处理完成
	wg.Wait()
	// 等待命令完成并回收进程
	return c.Wait()
}
