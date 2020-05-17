package installer

import (
	"os"
	"os/exec"
	"time"

	"golang.org/x/sys/windows/registry"
)

func Install() error {
	_, err := os.Stat("C:/Program Files/Krypton")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("C:/Program Files/Krypton", os.ModeDir)
			if err != nil {
				return err
			}
		}
	}

	time.Sleep(5 * time.Second)
	cleanup()
	os.Mkdir("C:/Program Files/Krypton/Settings", os.ModeDir)
	os.Mkdir("C:/Program Files/Krypton/Updates", os.ModeDir)

	exe, err := os.Executable()
	if err != nil {
		return err
	}
	err = copyFile(exe, "C:/Program Files/Krypton/Krypton.exe")
	if err != nil {
		return err
	}

	err = createRegistryKey()
	if err != nil {
		return err
	}
	err = createScheduledTasks()
	if err != nil {
		return err
	}
	return nil
}

func cleanup() {
	os.Remove("C:/Program Files/Krypton/7z.exe")
	os.Remove("C:/Program Files/Krypton/data.zip")
	os.Remove("C:/Program Files/Krypton/update.zip")
	os.Remove("C:/Program Files/Krypton/Settings.xml")
	os.RemoveAll("C:/Program Files/Krypton/data")
	os.RemoveAll("C:/Program Files/Krypton/update")
}

func createRegistryKey() error {
	_, _, err := registry.CreateKey(registry.LOCAL_MACHINE,
		`SOFTWARE\\Krypton`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	return nil
}

func createScheduledTasks() error {
	path := "C:/Program Files/Krypton/Krypton.exe"
	cmd := exec.Command("schtasks.exe", "/Create", "/SC", "HOURLY", "/TN",
		"KryptonUpdate", "/RU", "SYSTEM", "/F", "/TR", path+" --update")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("schtasks.exe", "/Create", "/SC", "DAILY", "/TN",
		"KryptonUpgrade", "/RU", "SYSTEM", "/F", "/TR", path+" --upgrade")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
