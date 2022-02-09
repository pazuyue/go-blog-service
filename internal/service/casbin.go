package service

type CasbinInfo struct {
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
}
type CasbinCreateRequest struct {
	RoleId      string       `json:"role_id" form:"role_id" description:"角色ID"`
	CasbinInfos []CasbinInfo `json:"casbin_infos" description:"权限模型列表"`
}

type CasbinListResponse struct {
	List []CasbinInfo `json:"list" form:"list"`
}

type CasbinListRequest struct {
	RoleID string `json:"role_id" form:"role_id"`
}

func (svc *Service) CasbinCreate(param *CasbinCreateRequest) error {
	for _, v := range param.CasbinInfos {
		err := svc.dao.CasbinCreate(param.RoleId, v.Path, v.Method)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) CasbinList(param *CasbinListRequest) [][]string {
	return s.dao.CasbinList(param.RoleID)
}
