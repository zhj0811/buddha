package define

import "github.com/pkg/errors"

const (
	Success = iota //成功
	QueryErr
	ReadRequestBodyErr
	PostTxErr
	UnmarshalErr
	MarshalErr
	MarshalContractArgsErr
	PreExecErr
	PreInvokeWasmErr

	InsertDBErr
	QueryDBErr
	UpdateDBErr
	DeleteDBErr

	QueryContractErr
	ContractRequestParamErr

	LoadFileErr
	IpfsAddErr
	IpfsCatErr

	ReaderErr
	ConvertErr
	RightErr
	RequestErr
	UnknownContractMethod

	ParamErr         //参数错误
	TokenErr         //验证t错误
	RegisterTokenErr //生成token错误
)

var (
	ErrRight        = errors.New("no permission")
	ErrRequest      = errors.New("request error")
	ErrContractTx   = errors.New("invalid contract tx")
	ErrContractArgs = errors.New("unmarshal contract args failed")
	ErrParam        = errors.New("invalid param")
	ErrCount        = errors.New("invalid count")
	ErrToken        = errors.New("invalid token")
)
