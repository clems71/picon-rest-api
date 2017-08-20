package main

import "os"

func getEnvOr(envVar string, def string) string {
	x := os.Getenv(envVar)
	if len(x) == 0 {
		x = def
	}
	return x
}

func clamp(x float32) float32 {
	if x > 1.0 {
		return 1.0
	} else if x < -1.0 {
		return -1.0
	}
	return x
}
