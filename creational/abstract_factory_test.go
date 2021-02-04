package creational

import "testing"

func getMainAndDetail(factory DAOFactory) (string, string) {
	mainData := factory.CreateOrderMainDAO().SaveOrderMain()
	detailData := factory.CreateOrderDetailDAO().SaveOrderDetail()

	return mainData, detailData
}

func TestRDBDAOFactory(t *testing.T) {
	f := &RDBDAOFactory{}
	mainData, detailData := getMainAndDetail(f)

	if mainData != "RDB save order main." {
		t.Fatal("RDB save order main error.")
	}

	if detailData != "RDB save order detail." {
		t.Fatal("RDB save order detail error.")
	}
}

func TestXMLDAOFactory(t *testing.T) {
	f := &XMLDAOFactory{}
	mainData, detailData := getMainAndDetail(f)

	if mainData != "XML save order main." {
		t.Fatal("XML save order main error.")
	}

	if detailData != "XML save order detail." {
		t.Fatal("XML save order detail error.")
	}
}
