package singularity

// SingularityRequest contains a high level information of
//  for a single project or deployable item.
type SingularityRequest struct {
	ID                                              string            `json:"id"`
	Instances                                       int64             `json:"instances"`
	NumRetriesOnFailure                             int64             `json:"numRetriesOnFailure"`
	QuartzSchedule                                  string            `json:"quartzSchedule"`
	RequestType                                     string            `json:"requestType"`
	Schedule                                        string            `json:"schedule"`
	ScheduleType                                    string            `json:"scheduleType"`
	HideEvenNumberAcrossRacksHint                   bool              `json:"hideEventNumerAcrossRacksHint"`
	TaskExecutionTimeLimitMillis                    int               `json:"taskExecutionTimeLimitMills"`
	TaskLogErrorRegexCaseSensitive                  bool              `json:"taskLogErrorRegexCaseSensitive"`
	SkipHealthchecks                                bool              `json:"skipHealthchecks"`
	WaitAtLeastMillisAfterTaskFinishesForReschedule int               `json:"waitAtleastMillisAfterTaskFinishesForReschedule"`
	TaskPriorityLevel                               int               `json:"taksPriorityLevel"`
	RackAffinity                                    []string          `json:"RackAffinity"`
	MaxTasksPerOffer                                int               `json:"maxTasksPerOffer"`
	BounceAfterScale                                bool              `json:"bounceAfterScale"`
	RackSensitive                                   bool              `json:"rackSensitive"`
	AllowedSlaveAttributes                          map[string]string `json:"allowedSlaveAttributes"`
	Owners                                          []string          `json:"owners"`
	RequiredRole                                    string            `json:"requiredRole"`
	ScheduledExpectedRuntimeMillis                  int               `json:"scheduledExpectedRuntimeMillis"`
	RequiredSlaveAttributes                         map[string]string `json:"requiredSlaveAttributes"`
	LoadBalanced                                    bool              `json:"loadBalanced"`
	KillOldNonLongRunningTasksAfterMillis           int               `json:"killOldNonLongRunningTasksAfterMillis"`
	ScheduleTimeZone                                string            `json:"scheduledTimeZone"`
	AllowBounceToSameHost                           bool              `json:"allowBounceToSamehost"`
	TaskLogErrorRegex                               string            `json:"taskLogErrorRegex"`
}

type ActiveDeploy struct {
	DeployID  string `json:"deployId"`
	RequestID string `json:"requestId"`
	Timestamp int64  `json:"timestamp"`
}

// SingularityDeployState contains specific configuration or version
// of the running code for that deployable item
type SingularityDeployState struct {
	ActiveDeploy `json:"activeDeploy"`
	RequestID    string `json:"requestId"`
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

// Task contains JSON response of /api/requests/request/ID.
type Task struct {
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
			ClubID             string      `json:"CLUB_ID"`
			ClubName           string      `json:"CLUB_NAME"`
			DockerImageVersion interface{} `json:"DOCKER_IMAGE_VERSION"`
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
		RequestType         string `json:"requestType"`
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

// RequestService contains required HTTP body to create a Singularity Request with requestType
// SERVICE.
type RequestService struct {
	ID                                              string            `json:"id"`
	Instances                                       int64             `json:"instances"`
	RequestType                                     string            `json:"requestType"`
	HideEvenNumberAcrossRacksHint                   bool              `json:"hideEventNumerAcrossRacksHint"`
	TaskExecutionTimeLimitMillis                    int               `json:"taskExecutionTimeLimitMills"`
	TaskLogErrorRegexCaseSensitive                  bool              `json:"taskLogErrorRegexCaseSensitive"`
	SkipHealthchecks                                bool              `json:"skipHealthchecks"`
	WaitAtLeastMillisAfterTaskFinishesForReschedule int               `json:"waitAtleastMillisAfterTaskFinishesForReschedule"`
	TaskPriorityLevel                               int               `json:"taksPriorityLevel"`
	RackAffinity                                    []string          `json:"RackAffinity"`
	MaxTasksPerOffer                                int               `json:"maxTasksPerOffer"`
	BounceAfterScale                                bool              `json:"bounceAfterScale"`
	RackSensitive                                   bool              `json:"rackSensitive"`
	AllowedSlaveAttributes                          map[string]string `json:"allowedSlaveAttributes"`
	Owners                                          []string          `json:"owners"`
	RequiredRole                                    string            `json:"requiredRole"`
	ScheduledExpectedRuntimeMillis                  int               `json:"scheduledExpectedRuntimeMillis"`
	RequiredSlaveAttributes                         map[string]string `json:"requiredSlaveAttributes"`
	LoadBalanced                                    bool              `json:"loadBalanced"`
	ScheduleTimeZone                                string            `json:"scheduledTimeZone"`
	AllowBounceToSameHost                           bool              `json:"allowBounceToSamehost"`
	TaskLogErrorRegex                               string            `json:"taskLogErrorRegex"`
}

// RequestOnDemand contains required HTTP body to create a Singularity Request with requestType
// ON_DEMAND.
type RequestOnDemand struct {
	ID                                              string            `json:"id"`
	NumRetriesOnFailure                             int64             `json:"numRetriesOnFailure"`
	RequestType                                     string            `json:"requestType"`
	HideEvenNumberAcrossRacksHint                   bool              `json:"hideEventNumerAcrossRacksHint"`
	TaskExecutionTimeLimitMillis                    int               `json:"taskExecutionTimeLimitMills"`
	TaskLogErrorRegexCaseSensitive                  bool              `json:"taskLogErrorRegexCaseSensitive"`
	SkipHealthchecks                                bool              `json:"skipHealthchecks"`
	WaitAtLeastMillisAfterTaskFinishesForReschedule int               `json:"waitAtleastMillisAfterTaskFinishesForReschedule"`
	TaskPriorityLevel                               int               `json:"taksPriorityLevel"`
	RackAffinity                                    []string          `json:"RackAffinity"`
	MaxTasksPerOffer                                int               `json:"maxTasksPerOffer"`
	BounceAfterScale                                bool              `json:"bounceAfterScale"`
	RackSensitive                                   bool              `json:"rackSensitive"`
	AllowedSlaveAttributes                          map[string]string `json:"allowedSlaveAttributes"`
	Owners                                          []string          `json:"owners"`
	RequiredRole                                    string            `json:"requiredRole"`
	ScheduledExpectedRuntimeMillis                  int               `json:"scheduledExpectedRuntimeMillis"`
	RequiredSlaveAttributes                         map[string]string `json:"requiredSlaveAttributes"`
	LoadBalanced                                    bool              `json:"loadBalanced"`
	KillOldNonLongRunningTasksAfterMillis           int               `json:"killOldNonLongRunningTasksAfterMillis"`
	ScheduleTimeZone                                string            `json:"scheduledTimeZone"`
	AllowBounceToSameHost                           bool              `json:"allowBounceToSamehost"`
	TaskLogErrorRegex                               string            `json:"taskLogErrorRegex"`
}

// RequestScheduled contains required HTTP body to create a Singularity Request with requestType
// Scheduled.
type RequestScheduled struct {
	ID                                              string            `json:"id"`
	NumRetriesOnFailure                             int64             `json:"numRetriesOnFailure"`
	RequestType                                     string            `json:"requestType"`
	Schedule                                        string            `json:"schedule"`
	HideEvenNumberAcrossRacksHint                   bool              `json:"hideEventNumerAcrossRacksHint"`
	TaskExecutionTimeLimitMillis                    int               `json:"taskExecutionTimeLimitMills"`
	TaskLogErrorRegexCaseSensitive                  bool              `json:"taskLogErrorRegexCaseSensitive"`
	SkipHealthchecks                                bool              `json:"skipHealthchecks"`
	WaitAtLeastMillisAfterTaskFinishesForReschedule int               `json:"waitAtleastMillisAfterTaskFinishesForReschedule"`
	TaskPriorityLevel                               int               `json:"taksPriorityLevel"`
	RackAffinity                                    []string          `json:"RackAffinity"`
	MaxTasksPerOffer                                int               `json:"maxTasksPerOffer"`
	BounceAfterScale                                bool              `json:"bounceAfterScale"`
	RackSensitive                                   bool              `json:"rackSensitive"`
	AllowedSlaveAttributes                          map[string]string `json:"allowedSlaveAttributes"`
	Owners                                          []string          `json:"owners"`
	RequiredRole                                    string            `json:"requiredRole"`
	ScheduledExpectedRuntimeMillis                  int               `json:"scheduledExpectedRuntimeMillis"`
	RequiredSlaveAttributes                         map[string]string `json:"requiredSlaveAttributes"`
	LoadBalanced                                    bool              `json:"loadBalanced"`
	KillOldNonLongRunningTasksAfterMillis           int               `json:"killOldNonLongRunningTasksAfterMillis"`
	ScheduleTimeZone                                string            `json:"scheduledTimeZone"`
	AllowBounceToSameHost                           bool              `json:"allowBounceToSamehost"`
	TaskLogErrorRegex                               string            `json:"taskLogErrorRegex"`
}

// RequestWorker contains required HTTP body to create a Singularity Request with requestType
// WORKER.
type RequestWorker struct {
	ID                                              string            `json:"id"`
	Instances                                       int64             `json:"instances"`
	RequestType                                     string            `json:"requestType"`
	HideEvenNumberAcrossRacksHint                   bool              `json:"hideEventNumerAcrossRacksHint"`
	TaskExecutionTimeLimitMillis                    int               `json:"taskExecutionTimeLimitMills"`
	TaskLogErrorRegexCaseSensitive                  bool              `json:"taskLogErrorRegexCaseSensitive"`
	SkipHealthchecks                                bool              `json:"skipHealthchecks"`
	WaitAtLeastMillisAfterTaskFinishesForReschedule int               `json:"waitAtleastMillisAfterTaskFinishesForReschedule"`
	TaskPriorityLevel                               int               `json:"taksPriorityLevel"`
	RackAffinity                                    []string          `json:"RackAffinity"`
	MaxTasksPerOffer                                int               `json:"maxTasksPerOffer"`
	BounceAfterScale                                bool              `json:"bounceAfterScale"`
	RackSensitive                                   bool              `json:"rackSensitive"`
	AllowedSlaveAttributes                          map[string]string `json:"allowedSlaveAttributes"`
	Owners                                          []string          `json:"owners"`
	RequiredRole                                    string            `json:"requiredRole"`
	ScheduledExpectedRuntimeMillis                  int               `json:"scheduledExpectedRuntimeMillis"`
	RequiredSlaveAttributes                         map[string]string `json:"requiredSlaveAttributes"`
	LoadBalanced                                    bool              `json:"loadBalanced"`
	ScheduleTimeZone                                string            `json:"scheduledTimeZone"`
	AllowBounceToSameHost                           bool              `json:"allowBounceToSamehost"`
	TaskLogErrorRegex                               string            `json:"taskLogErrorRegex"`
}

// RequestRunOnce contains required HTTP body to create a Singularity Request with requestType
// RUN_ONCE.
type RequestRunOnce struct {
	ID                                              string            `json:"id"`
	Instances                                       int64             `json:"instances"`
	RequestType                                     string            `json:"requestType"`
	HideEvenNumberAcrossRacksHint                   bool              `json:"hideEventNumerAcrossRacksHint"`
	TaskExecutionTimeLimitMillis                    int               `json:"taskExecutionTimeLimitMills"`
	TaskLogErrorRegexCaseSensitive                  bool              `json:"taskLogErrorRegexCaseSensitive"`
	SkipHealthchecks                                bool              `json:"skipHealthchecks"`
	WaitAtLeastMillisAfterTaskFinishesForReschedule int               `json:"waitAtleastMillisAfterTaskFinishesForReschedule"`
	TaskPriorityLevel                               int               `json:"taksPriorityLevel"`
	RackAffinity                                    []string          `json:"RackAffinity"`
	MaxTasksPerOffer                                int               `json:"maxTasksPerOffer"`
	BounceAfterScale                                bool              `json:"bounceAfterScale"`
	RackSensitive                                   bool              `json:"rackSensitive"`
	AllowedSlaveAttributes                          map[string]string `json:"allowedSlaveAttributes"`
	Owners                                          []string          `json:"owners"`
	RequiredRole                                    string            `json:"requiredRole"`
	ScheduledExpectedRuntimeMillis                  int               `json:"scheduledExpectedRuntimeMillis"`
	RequiredSlaveAttributes                         map[string]string `json:"requiredSlaveAttributes"`
	LoadBalanced                                    bool              `json:"loadBalanced"`
	ScheduleTimeZone                                string            `json:"scheduledTimeZone"`
	AllowBounceToSameHost                           bool              `json:"allowBounceToSamehost"`
	TaskLogErrorRegex                               string            `json:"taskLogErrorRegex"`
}

type RunNowRequest struct {
}

// DeleteRequest contains HTTP body for a Delete Singularity Request. Please see below URL for
// more information.
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#delete-apirequestsrequestrequestid
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#-singularitydeleterequestrequest
type DeleteRequest struct {
	DeleteFromLoadBalancer bool   `json:"deleteFromLoadBalancer"` //optional	Should the service associated with the request be removed from the load balancer
	Message                string `json:"message"`                //optional	A message to show to users about why this action was taken
	ActionID               string `json:"actionId"`               //An id to associate with this action for metadata purposes
}
