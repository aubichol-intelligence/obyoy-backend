package apipattern

// UserList holds the api for giving user list
const UserList string = "/api/v1/alluser"

// LoginUser holds the api for logging user
const LoginUser string = "/api/v1/users/login"

// LoginUser holds the api for logging user
const LogoutUser string = "/api/v1/users/logout"

// RegistrationToken holds holds the api for giving registration token
const RegistrationToken string = "/api/v1/registration/token"

// UserRegistration holds the api for registering a user
const UserRegistration string = "/api/v1/users/registration"

// UserSearch holds the api for searching user
const UserSearch string = "/api/v1/users/search"

// DatasetCreate holds the api string for creating a dataset
const DatasetCreate string = "/api/v1/dataset/create"

// DatasetRead holds the api string for reading datasets
const DatasetRead string = "/api/v1/dataset/get/{id}"

// DatasetUpdate holds the api string for updating dataset
const DatasetUpdate string = "/api/v1/dataset/update"

// DatasetDelete holds the api string for getting a dataset
const DatasetDelete string = "/api/v1/dataset/delete"

// DatasetList holds the api string for getting a dataset list
const DatesetList string = "/api/v1/dataset/list/{skip}/{limit}"

// DatastreamCreate holds the api string for creating a datastream
const DatastreamCreate string = "/api/v1/datastream/create"

// DatastreamRead holds the api string for reading datastreams
const DatastreamRead string = "/api/v1/datastream/get"

// DatastreamRead holds the api string for reading datastreams
const DatastreamReadNext string = "/api/v1/datastream/getnext"

// DatastreamUpdate holds the api string for updating datastream
const DatastreamUpdate string = "/api/v1/datastream/update"

// DatastreamDelete holds the api string for getting a datastream
const DatastreamDelete string = "/api/v1/datastream/delete"

// ParallelsentenceCreate holds the api string for creating a parallelsentence
const ParallelsentenceCreate string = "/api/v1/parallelsentence/create"

// ParallelsentenceRead holds the api string for reading parallelsentences
const ParallelsentenceRead string = "/api/v1/parallelsentence/get/{id}"

// ParallelsentenceUpdate holds the api string for updating parallelsentence
const ParallelsentenceUpdate string = "/api/v1/parallelsentence/update"

// ParallelsentenceDelete holds the api string for getting a parallelsentence
const ParallelsentenceDelete string = "/api/v1/parallelsentence/delete"

// ParallelsentenceList holds the api string for getting a dataset list
const ParallelsentenceList string = "/api/v1/parallelsentence/list/{skip}/{limit}"

// TranslationCreate holds the api string for creating a translation
const TranslationCreate string = "/api/v1/translation/create"

// TranslationRead holds the api string for reading translations
const TranslationRead string = "/api/v1/translation/get/{id}"

// TranslationUpdate holds the api string for updating translation
const TranslationUpdate string = "/api/v1/translation/update"

// TranslationDelete holds the api string for getting a translation
const TranslationDelete string = "/api/v1/translation/delete"
