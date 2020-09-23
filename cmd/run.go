/*
Copyright © 2020 Elias Hussary <eliashussary@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/eliashussary/wz-assist/assist"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func toStringMapInt(m map[string]string) map[string]int64 {
	mp := make(map[string]int64)
	for k, v := range m {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			i = 0
		}
		mp[k] = i
	}
	return mp
}

func toMilliseconds(s interface{}) time.Duration {
	v, _ := s.(string)
	d, _ := time.ParseDuration(v)
	return d
}

func toInt(s interface{}) int {
	v, _ := s.(int)
	return v
}
func toUint16(s interface{}) uint16 {

	i, _ := s.(int)
	v := uint16(i)

	return v
}

func printMaps(maps ...map[string]interface{}) {
	for _, m := range maps {
		for k, v := range m {
			fmt.Println(k, "=", v)
		}
	}
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the assist",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		keyBindsConfig := viper.GetStringMap("keybinds")
		delaysConfig := viper.GetStringMap("delays")
		recoilSenseConfig := viper.GetStringMap("recoilsense")

		app := assist.NewAssist(
			assist.Keybinds{
				ADS:               toInt(keyBindsConfig["ads"]),
				Ping:              toUint16(keyBindsConfig["ping"]),
				RecoilToggle:      toInt(keyBindsConfig["recoiltoggle"]),
				RapidFireToggle:   toInt(keyBindsConfig["rapidfiretoggle"]),
				RecoilSenseToggle: toInt(keyBindsConfig["recoilsensetoggle"]),
			},
			assist.Delays{
				Standard:       toMilliseconds(delaysConfig["standard"]),
				AutoPingDelay:  toMilliseconds(delaysConfig["autopingdelay"]),
				RapidFireDelay: toMilliseconds(delaysConfig["rapidfiredelay"]),
			},
			assist.RecoilSense{
				Low:  toMilliseconds(recoilSenseConfig["low"]),
				High: toMilliseconds(recoilSenseConfig["high"]),
			},
		)

		app.Start()
		fmt.Println("Started...")
		fmt.Printf("\nConfig\n============\n")
		printMaps(keyBindsConfig, delaysConfig, recoilSenseConfig)
		select {}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
