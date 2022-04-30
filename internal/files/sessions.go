package files

func Sessions() ([]string, error) {
	var sessions []string

	err := IterateSessionDirs("", func(sessionName string) error {
		sessions = append(sessions, sessionName)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func SessionUrls(session string) ([]string, error) {
	var urls []string
	err := IterateStatusFiles(session, func(statusFile PageDownloadStatus) error {
		urls = append(urls, statusFile.Request.URL.String())

		return nil
	})
	if err != nil {
		return nil, err
	}

	return urls, nil
}
