// 抽象工厂模式用于生成产品族的工厂，所生成的对象是有关联的
// 如果抽象工厂退化成生成的对象无关联则成为工厂函数模式

package creational

type OrderMainDAO interface {
	SaveOrderMain() string
}

type OrderDetailDAO interface {
	SaveOrderDetail() string
}

type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type RDBMainDAO struct {
}

func (*RDBMainDAO) SaveOrderMain() string {
	return "RDB save order main."
}

type RDBDetailDAO struct {
}

func (*RDBDetailDAO) SaveOrderDetail() string {
	return "RDB save order detail."
}

type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

type XMLMainDAO struct {
}

func (*XMLMainDAO) SaveOrderMain() string {
	return "XML save order main."
}

type XMLDetailDAO struct {
}

func (*XMLDetailDAO) SaveOrderDetail() string {
	return "XML save order detail."
}

type XMLDAOFactory struct {
}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
