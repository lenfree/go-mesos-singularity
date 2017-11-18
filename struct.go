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

// ActiveDeploy have a string deployId, requestId and a timestamp.
type ActiveDeploy struct {
	DeployID  string `json:"deployId"`
	RequestID string `json:"requestId"`
	Timestamp int64  `json:"timestamp"`
}

// RequestDeployState contains specific configuration or version
// of the running code for that deployable item
type RequestDeployState struct {
	ActiveDeploy `json:"activeDeploy"`
	RequestID    string `json:"requestId"`
}

// Request struct contains all singularity requests.
// This have a JSON response of /api/requests.
type Request struct {
	SingularityRequest `json:"request"`
	RequestDeployState `json:"requestDeployState"`
	State              string `json:"state"`
}

// Requests is a slice of Request.
type Requests []Request

// ContainerInfo contains information about a Docker type Singularity
// container type.
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#-singularitycontainerinfo
type ContainerInfo struct {
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
}

// Task contains JSON response of /api/requests/request/ID.
type Task struct {
	ActiveDeploy struct {
		Arguments     []string `json:"arguments"`
		Command       string   `json:"command"`
		ContainerInfo `json:"containerInfo"`
		Env           struct {
			ClubID             string      `json:"CLUB_ID"`
			ClubName           string      `json:"CLUB_NAME"`
			DockerImageVersion interface{} `json:"DOCKER_IMAGE_VERSION"`
		} `json:"env"`
		ID                         string `json:"id"`
		RequestID                  string `json:"requestId"`
		SingularityDeployResources `json:"resources"`
		Uris                       []string `json:"uris"`
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

// SingularityDeployResources includes information about required/configured
// resources needed for a request/job.
type SingularityDeployResources struct {
	Cpus     float64 `json:"cpus"`
	MemoryMb float64 `json:"memoryMb"`
	NumPorts int64   `json:"numPorts"`
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

// SingularityScaleRequest contains parameters for making scaling a request. For more info, please see:
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#-singularityscalerequest
type SingularityScaleRequest struct {
	SkipHealthchecks bool   `json:"skipHealthchecks"`
	DurationMillis   int64  `json:"durationMillis"`
	Bounce           bool   `json:"bounce"`
	Message          string `json:"message"`
	ActionID         string `json:"actionId"`
	Instances        int    `json:"instances"`
	Incremental      int    `json:"incremental"`
}

// SingularityExpiringSkipHealthchecks have parameters for a expiring skip
// healthchecks.
type SingularityExpiringSkipHealthchecks struct {
	User                                string      `json:"user"`
	RequestID                           string      `json:"requestId"`
	StartMillis                         int64       `json:"startMillis"`
	ActionID                            string      `json:"actionId"`
	SingularityExpiringAPIRequestObject interface{} `json:"expiringAPIRequestObject"`
	RevertToSkipHealthchecks            bool        `json:"revertToSkipHealthchecks"`
}

// RequestState contains a string state of a existing job/Singulariy request. Allowable
// values are:  ACTIVE, DELETING, DELETED, PAUSED, SYSTEM_COOLDOWN, FINISHED,
// DEPLOYING_TO_UNPAUSE
type RequestState struct {
	State string `json:"state"` //Allowable values:
}

// HealthcheckProtocol contains a string with allowable value of
// HTTP or HTTPS.
type HealthcheckProtocol string

// HealthcheckOptions contains parameters of a healthcheck options
// for a new and existing Singularity request.
type HealthcheckOptions struct {
	StartupDelaySeconds    int    `json:"startupDelaySeconds"`
	ResponseTimeoutSeconds int    `json:"responseTimeoutSeconds"`
	IntervalSeconds        int    `json:"intervalSeconds"`
	URI                    string `json:"uri"` //Healthcheck uri to hit
	FailureStatusCodes     []int  `json:"failureStatusCodes"`
	MaxRetries             int    `json:"maxRetries"`
	StartupTimeoutSeconds  int    `json:"startupTimeoutSeconds"`
	PortNumber             int    `json:"portNumber"`
	StartupIntervalSeconds int    `json:"startupIntervalSeconds"` //Time to wait after a failed healthcheck to try again in seconds.
	HealthcheckProtocol    `json:"protocol"`
	PortIndex              int `json:"portIndex"`
}

// SingularityDeploy contains requird and optional parameter to configure
// a new and existing Singularity deploy.
type SingularityDeploy struct {
	CustomExecutorID           string `json:"customExecutorId"`
	SingularityDeployResources `json:"resources"`
	Uris                       interface{} `json:"uris"` //Array[SingularityMesosArtifact]	optional	List of URIs to download before executing the deploy command.
	ContainerInfo              `json:"containerInfo"`
	// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#-set List of domains to host this service on, for use with the load balancer api
	LoadBalancerDomains    interface{} `json:"loadBalancerDomains"` //Set
	HealthcheckOptions     `json:"healthcheck"`
	Arguments              []string    `json:"arguments"`
	TaskEnv                interface{} `json:"taskEnv"` //Map[int,Map[string,string]]	Map of environment variable overrides for specific task instances.
	AutoAdvanceDeploySteps bool        `json:"autoAdvanceDeploySteps"`
}

// SingularityExpiringPause contains information of a existing
// Singularity request.
type SingularityExpiringPause struct {
	User                                string      `json:"user"`
	RequestID                           string      `json:"requestId"`
	StartMillis                         int64       `json:"startMillis"`
	ActionID                            string      `json:"actionId"`
	SingularityExpiringAPIRequestObject interface{} `json:"expiringAPIRequestObject"`
}

// SingularityExpiringBounce contains information of a existing
// Singularity request.
type SingularityExpiringBounce struct {
	User                     string      `json:"user"`
	RequestID                string      `json:"requestId"`
	StartMillis              int64       `json:"startMillis"`
	DeployID                 string      `json:"deployId"`
	ActionID                 string      `json:"actionId"`
	ExpiringAPIRequestObject interface{} `json:"expiringAPIRequestObject"`
}

// SingularityDeployProgress contains deploy progress of a existing
// Singularity request.
type SingularityDeployProgress struct {
	AutoAdvanceDeploySteps     bool        `json:"autoAdvanceDeploySteps"`
	StepComplete               bool        `json:"stepComplete"`
	DeployStepWaitTimeMs       int64       `json:"deployStepWaitTimeMs"`
	Timestamp                  int64       `json:"timestamp"`
	DeployInstanceCountPerStep int         `json:"deployInstanceCountPerStep"`
	FailedDeployTasks          interface{} `json:"failedDeployTasks"` //Set	optional
	CurrentActiveInstances     int         `json:"currentActiveInstances"`
	TargetActiveInstances      int         `json:"targetActiveInstances"`
}

// SingularityLoadBalancerRequestID have loadbalancer information of a
// Singularity request.
type SingularityLoadBalancerRequestID struct {
	//optional	Allowable values: ADD, REMOVE, DEPLOY, DELETE
	RequestType   string `json:"requestType"`
	AttemptNumber int    `json:"attemptNumber"`
	ID            string `json:"id"`
}

// SingularityLoadBalancerUpdate contains parameters required to update a
// Singularity request's loadbalancer.
type SingularityLoadBalancerUpdate struct {
	// Allowable values: UNKNOWN, FAILED, WAITING, SUCCESS, CANCELING, CANCELED, INVALID_REQUEST_NOOP
	LoadBalancerState                string `json:"loadBalancerState"`
	SingularityLoadBalancerRequestID `json:"loadBalancerRequestId"`
	URI                              string `json:"uri"`
	// Allowable values: PRE_ENQUEUE, ENQUEUE, CHECK_STATE, CANCEL, DELETE
	Method    string `json:"method"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// SingularityDeployMarker holds information of a Singularity deploy.
type SingularityDeployMarker struct {
	User      string `json:"user"`
	RequestID string `json:"requestId"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	DeployID  string `json:"deployId"`
}

// SingularityDeployState holds information of a existing Singularity
// deploy.
type SingularityDeployState struct {
	// Allowable values: SUCCEEDED, FAILED_INTERNAL_STATE, CANCELING, WAITING, OVERDUE, FAILED, CANCELED
	CurrentDeployState            string `json:"currentDeployState"`
	SingularityRequest            `json:"updatedRequest"`
	SingularityDeployProgress     `json:"deployProgress"`
	SingularityLoadBalancerUpdate `json:"lastLoadBalancerUpdate"`
	SingularityDeployMarker       `json:"deployMarker"`
}

type SingularityExpiringAPIRequestObject struct {
	ActionID         string `json:"actionId"`
	DurationMillis   int64  `json:"durationMillis"`
	Instances        int64  `json:"instances"`
	Message          string `json:"message"`
	SkipHealthchecks bool   `json:"skipHealthchecks"`
}

// SingularityExpiringScale holds information of a expiring scale Singularity
// deploy.
type SingularityExpiringScale struct {
	RevertToInstances                   int    `json:"revertToInstances"`
	User                                string `json:"user"`
	RequestID                           string `json:"requestId"`
	Bounce                              bool   `json:"bounce"`
	StartMillis                         int64  `json:"startMillis"`
	ActionID                            string `json:"actionId"`
	DurationMillis                      int64  `json:"durationMillis"`
	SingularityExpiringAPIRequestObject `json:"expiringAPIRequestObject"`
}

// SingularityPendingDeploy holds information of a pending Singularity
// deploy.
type SingularityPendingDeploy struct {
	CurrentDeployState            string `json:"currentDeployState"` // Allowable values: SUCCEEDED, FAILED_INTERNAL_STATE, CANCELING, WAITING, OVERDUE, FAILED, CANCELED
	SingularityRequest            `json:"updatedRequest"`
	SingularityDeployProgress     `json:"deployProgress"`
	SingularityLoadBalancerUpdate `json:"lastLoadBalancerUpdate"`
	SingularityDeployMarker       `json:"deployMarker"`
}

// SingularityRequestParent contains request of a Singularity deploy.
type SingularityRequestParent struct {
	SingularityExpiringSkipHealthchecks `json:"expiringSkipHealthchecks"`
	PendingDeploy                       struct {
		CustomExecutorID           string `json:"customExecutorId"`
		SingularityDeployResources `json:"resources"`
		Uris                       interface{} `json:"uris"` //Array[SingularityMesosArtifact]	optional	List of URIs to download before executing the deploy command.
		ContainerInfo              `json:"containerInfo"`
		// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#-set List of domains to host this service on, for use with the load balancer api
		LoadBalancerDomains    interface{} `json:"loadBalancerDomains"` //Set
		HealthcheckOptions     `json:"healthcheck"`
		Arguments              []string    `json:"arguments"`
		TaskEnv                interface{} `json:"taskEnv"` //Map[int,Map[string,string]]	Map of environment variable overrides for specific task instances.
		AutoAdvanceDeploySteps bool        `json:"autoAdvanceDeploySteps"`
	} `json:"pendingDeploy"`
	ActiveDeploy struct {
		CustomExecutorID           string `json:"customExecutorId"`
		SingularityDeployResources `json:"resources"`
		Uris                       interface{} `json:"uris"` //Array[SingularityMesosArtifact]	optional	List of URIs to download before executing the deploy command.
		ContainerInfo              `json:"containerInfo"`
		// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#-set List of domains to host this service on, for use with the load balancer api
		LoadBalancerDomains    interface{} `json:"loadBalancerDomains"` //Set
		HealthcheckOptions     `json:"healthcheck"`
		Arguments              []string    `json:"arguments"`
		TaskEnv                interface{} `json:"taskEnv"` //Map[int,Map[string,string]]	Map of environment variable overrides for specific task instances.
		AutoAdvanceDeploySteps bool        `json:"autoAdvanceDeploySteps"`
	} `json:"activeDeploy"`
	SingularityExpiringPause  `json:"expiringPause"`
	SingularityExpiringBounce `json:"expiringBounce"`
	SingularityRequest        `json:"request"`
	SingularityPendingDeploy  `json:"pendingDeployState"`
	SingularityExpiringScale  `json:"expiringScale"`
	SingularityDeployState    `json:"requestDeployState"`
	State                     string `json:"state"`
}

type RunNowRequest struct {
}

// SingularityDeleteRequest contains HTTP body for a Delete Singularity Request. Please see below URL for
// more information.
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#delete-apirequestsrequestrequestid
// https://github.com/HubSpot/Singularity/blob/master/Docs/reference/api.md#-singularitydeleterequestrequest
type SingularityDeleteRequest struct {
	DeleteFromLoadBalancer bool   `json:"deleteFromLoadBalancer"` //optional	Should the service associated with the request be removed from the load balancer
	Message                string `json:"message"`                //optional	A message to show to users about why this action was taken
	ActionID               string `json:"actionId"`               //An id to associate with this action for metadata purposes
}
