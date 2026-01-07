package main

import "fmt"

func banner() {
	lines := []string{
		"::::::::::: :::::::::  :::::::::   ::::::::  ",
		"    :+:     :+:    :+: :+:    :+: :+:    :+: ",
		"    +:+     +:+    +:+ +:+    +:+ +:+    +:+ ",
		"    +#+     +#++:++#+  +#+    +:+ +#+    +#+ ",
		"    +#+     +#+        +#+    +:+ +#+    +#+ ",
		"    #+#     #+#        #+#    #+# #+#    #+# ",
		"########### ###        #########   ######## ",
		"",
	}

	for i, line := range lines {
		color := ""
		switch i {
		case 0, 1, 2:
			color = "\x1b[95m"
		case 3, 4, 5: // orta mor
			color = "\x1b[35m" 
		case 6, 7, 8, 9:
			color = "\x1b[90m" 
		}
		fmt.Println(color + line + "\x1b[0m") 
	}
}
