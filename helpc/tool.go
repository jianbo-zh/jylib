package helpc

import "fmt"

func ServerName(env string, name string) string {
	return fmt.Sprintf("%s.%s", env, name)
}

func ServerEndpoint(env string, name string) string {
	return "discovery:///" + ServerName(env, name)
}
