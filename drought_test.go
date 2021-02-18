package main

import (
	"planets/utils"
	"testing"
)

func TestGradesBetasoide(t *testing.T) {
	Betasoide90 := 3 * 90
	Betasoide180 := 3 * 180
	Betasoide270 := 3 * 270
	Betasoide360 := 3 * 360
	if !utils.ValidateGrade(0) {
		t.Error("Error to validate position Betasoide to 0 days")
	}
	if !utils.ValidateGrade(Betasoide90) {
		t.Error("Error to validate position Betasoide to 90 days")
	}
	if !utils.ValidateGrade(Betasoide180) {
		t.Error("Error to validate position Betasoide to 180 days")
	}
	if !utils.ValidateGrade(Betasoide270) {
		t.Error("Error to validate position Betasoide to 270 days")
	}
	if !utils.ValidateGrade(Betasoide360) {
		t.Error("Error to validate position Betasoide to 360 days")
	}
}

func TestGradesFarengi(t *testing.T) {
	Farengi90 := 1 * 90
	Farengi180 := 1 * 180
	Farengi270 := 1 * 270
	Farengi360 := 1 * 360
	if !utils.ValidateGrade(0) {
		t.Error("Error to validate position Farengi to 0 days")
	}
	if !utils.ValidateGrade(Farengi90) {
		t.Error("Error to validate position Farengi to 90 days")
	}
	if !utils.ValidateGrade(Farengi180) {
		t.Error("Error to validate position Farengi to 180 days")
	}
	if !utils.ValidateGrade(Farengi270) {
		t.Error("Error to validate position Farengi to 270 days")
	}
	if !utils.ValidateGrade(Farengi360) {
		t.Error("Error to validate position Farengi to 360 days")
	}
}

func TestGradesVulcano(t *testing.T) {
	Vulcano90 := 5 * 90
	Vulcano180 := 5 * 180
	Vulcano270 := 5 * 270
	Vulcano360 := 5 * 360
	if !utils.ValidateGrade(0) {
		t.Error("Error to validate position Vulcano to 0 days")
	}
	if !utils.ValidateGrade(Vulcano90) {
		t.Error("Error to validate position Vulcano to 90 days")
	}
	if !utils.ValidateGrade(Vulcano180) {
		t.Error("Error to validate position Vulcano to 180 days")
	}
	if !utils.ValidateGrade(Vulcano270) {
		t.Error("Error to validate position Vulcano to 270 days")
	}
	if !utils.ValidateGrade(Vulcano360) {
		t.Error("Error to validate position Vulcano to 360 days")
	}
}

func TestDrought90Grades(t *testing.T) {
	Betasoide90 := 3 * 90
	Farengi90 := 1 * 90
	Vulcano90 := 5 * 90

	if !utils.ValidateGrade(Betasoide90) && !utils.ValidateGrade(Farengi90) && !utils.ValidateGrade(Vulcano90) {
		t.Error("Test Fail when the planes not align to 90 days")
	}
}
func TestDrought180Grades(t *testing.T) {
	Betasoide180 := 3 * 180
	Farengi180 := 1 * 180
	Vulcano180 := 5 * 180

	if !utils.ValidateGrade(Betasoide180) && !utils.ValidateGrade(Farengi180) && !utils.ValidateGrade(Vulcano180) {
		t.Error("Test Fail when the planes not align to 180 days")
	}
}
func TestDrought270Grades(t *testing.T) {
	Betasoide270 := 3 * 270
	Farengi270 := 1 * 270
	Vulcano270 := 5 * 270

	if !utils.ValidateGrade(Betasoide270) && !utils.ValidateGrade(Farengi270) && !utils.ValidateGrade(Vulcano270) {
		t.Error("Test Fail when the planes not align to 270 days")
	}
}
func TestDrought360Grades(t *testing.T) {
	Betasoide360 := 3 * 360
	Farengi360 := 1 * 360
	Vulcano360 := 5 * 360

	if !utils.ValidateGrade(Betasoide360) && !utils.ValidateGrade(Farengi360) && !utils.ValidateGrade(Vulcano360) {
		t.Error("Test Fail when the planes not align to 360 days")
	}
}
