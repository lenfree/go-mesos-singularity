package singularity

// SingularityRequest contains a high level information of
//  for a single project or deployable item.
type SingularityRequest struct {
	ID                  string `json:"id"`
	Instances           int64  `json:"instances"`
	NumRetriesOnFailure int64  `json:"numRetriesOnFailure"`
	QuartzSchedule      string `json:"quartzSchedule"`
	RequestType         string `json:"requestType"`
	Schedule            string `json:"schedule"`
	ScheduleType        string `json:"scheduleType"`
}

// SingularityDeployState contains specific configuration or version
// of the running code for that deployable item
type SingularityDeployState struct {
	ActiveDeploy struct {
		DeployID  string `json:"deployId"`
		RequestID string `json:"requestId"`
		Timestamp int64  `json:"timestamp"`
	} `json:"activeDeploy"`
	RequestID string `json:"requestId"`
}

// Request struct contains all singularity requests.
// This have a JSON response of /api/requests.
type Request struct {
	SingularityRequest     `json:"request"`
	SingularityDeployState `json:"requestDeployState"`
	State                  string `json:"state"`
}

// Requests is a slice of Request.
type Requests []Request

// RequestDockerID contains JSON response of /api/requests/request/ID.
type RequestDockerID struct {
	ActiveDeploy struct {
		Arguments     []string `json:"arguments"`
		Command       string   `json:"command"`
		ContainerInfo struct {
			Docker struct {
				ForcePullImage bool     `json:"forcePullImage"`
				Image          string   `json:"image"`
				Network        string   `json:"network"`
				Parameters     struct{} `json:"parameters"`
				PortMappings   []struct {
					ContainerPort     int64  `json:"containerPort"`
					ContainerPortType string `json:"containerPortType"`
					HostPort          int64  `json:"hostPort"`
					HostPortType      string `json:"hostPortType"`
					Protocol          string `json:"protocol"`
				} `json:"portMappings"`
				Privileged bool `json:"privileged"`
			} `json:"docker"`
			Type    string `json:"type"`
			Volumes []struct {
				ContainerPath string `json:"containerPath"`
				HostPath      string `json:"hostPath"`
				Mode          string `json:"mode"`
			} `json:"volumes"`
		} `json:"containerInfo"`
		Env struct {
			ClubID   string `json:"CLUB_ID"`
			ClubName string `json:"CLUB_NAME"`
		} `json:"env"`
		ID        string `json:"id"`
		RequestID string `json:"requestId"`
		Resources struct {
			Cpus     float64 `json:"cpus"`
			MemoryMb int64   `json:"memoryMb"`
			NumPorts int64   `json:"numPorts"`
		} `json:"resources"`
		Uris []string `json:"uris"`
	} `json:"activeDeploy"`
	Request struct {
		ID                  string `json:"id"`
		NumRetriesOnFailure int64  `json:"numRetriesOnFailure"`
		QuartzSchedule      string `json:"quartzSchedule"`
		RequestType         string `json:"requestType"`
		Schedule            string `json:"schedule"`
		ScheduleType        string `json:"scheduleType"`
	} `json:"request"`
	RequestDeployState struct {
		ActiveDeploy struct {
			DeployID  string `json:"deployId"`
			RequestID string `json:"requestId"`
			Timestamp int64  `json:"timestamp"`
		} `json:"activeDeploy"`
		RequestID string `json:"requestId"`
	} `json:"requestDeployState"`
	State string `json:"state"`
}

type Resources struct {
	NumPorts int
	MemoryMb int
	DiskMb   int
	CPUs     int
}

type RunNowRequest struct {
}
