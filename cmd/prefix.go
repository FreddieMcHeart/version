package cmd

var Prefix string

func init() {
	rootCmd.PersistentFlags().StringVarP(&Prefix, "prefix", "p", "", "Your version`s prefix")
}
