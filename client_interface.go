package conn

type (
	ConnClient interface {
		//2018-12-12
		RegisterFn(FnInfo, func(Session, frist Data)) error
		//2018-12-12
		CallFn(Call) (Session, error)
		// RegisterApp(AppInfo) (App, error)
	}

	Session interface {
		Recv() (Data, error)
		Send(Data) error
		End(Data) error
	}
)

type (
	FnInfo struct {
		Fn          string
		Description string
		AppFilter
		ClientFilter
	}

	Call struct {
		App       string
		Fn        string
		ChunkType string
		Chunk     []byte
		Take      int32
		ClientFilter
	}

	Data struct {
		ChunkType string
		Chunk     []byte
		Take      int32
	}
)

type AppFilter map[string]interface{}
type ClientFilter map[string]interface{}
