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

// DatasetCreate holds the api string for creating a dataset
const DatastreamCreate string = "/api/v1/datastream/create"

// DatasetRead holds the api string for reading datasets
const DatastreamRead string = "/api/v1/datastream/get"

// DatasetRead holds the api string for reading datasets
const DatastreamReadNext string = "/api/v1/datastream/getnext"

// DatasetUpdate holds the api string for updating dataset
const DatastreamUpdate string = "/api/v1/datastream/update"

// DatasetDelete holds the api string for getting a dataset
const DatastreamDelete string = "/api/v1/datastream/delete"

// DatasetCreate holds the api string for creating a dataset
const ParallelsentenceCreate string = "/api/v1/parallelsentence/create"

// DatasetRead holds the api string for reading datasets
const ParallelsentenceRead string = "/api/v1/parallelsentence/get/{id}"

// DatasetUpdate holds the api string for updating dataset
const ParallelsentenceUpdate string = "/api/v1/parallelsentence/update"

// DatasetDelete holds the api string for getting a dataset
const ParallelsentenceDelete string = "/api/v1/parallelsentence/delete"

// DatasetCreate holds the api string for creating a dataset
const TranslationCreate string = "/api/v1/translation/create"

// DatasetRead holds the api string for reading datasets
const TranslationRead string = "/api/v1/translation/get/{id}"

// DatasetUpdate holds the api string for updating dataset
const TranslationUpdate string = "/api/v1/translation/update"

// DatasetDelete holds the api string for getting a dataset
const TranslationDelete string = "/api/v1/translation/delete"
