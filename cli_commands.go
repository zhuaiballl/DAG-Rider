package DAG_Rider

func (cli *CLI) startNode() {
	nd := Node{}
	nd.Start()
	select {}
}
