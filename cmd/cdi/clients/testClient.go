package clients

type TestClient struct {
	branchCount int
}

func NewTestClient() *TestClient {
	return &TestClient{branchCount: 23}
}

func (c *TestClient) RawBranchCount(gitRepoLink string) (int, error) {
	return c.branchCount, nil
}