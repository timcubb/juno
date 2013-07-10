package storm

// started using: http://www.stormondemand.com/api/docs/v1/Storm/Server.html#method_clone  maybe switch to: http://www.stormondemand.com/api/docs/v1/Server.html#method_available???

// use" http://golang.org/pkg/net/http/#Request.SetBasicAuth
// and: http://www.stormondemand.com/api/docs/tutorials/client.html
// to get the go client working


type StormCloneRequest struct {
    config_id int       //   A positive integer value (i.e. 1 and up).   
    domain string       //A fully-qualified domain name (i.e. liquidweb.com, www.liquidweb.com, etc)  * Required
    ip_count int        //A positive integer value (i.e. 1 and up).   
    password string     // A password of at least 7 characters and up to 30 characters in length, containing no spaces. Must contain 3 of the 4 following classes: lowercase, uppercase, numbers and punctuation.  * Required
    uniq_id string      // A six-character identifier, containing only capital letters and digits. * Required
    zone int            // A non-negative integer value (i.e. 0 and up).
}

type StormCloneResponse struct {
    server StormServer
}


type StormCreateRequest struct {
    antivirus string        // A string.   
    backup_enabled bool     // A boolean value (0 or 1).   * Optional
    backup_id int           // A positive integer value (i.e. 1 and up).   
    backup_plan string      // A single word, consisting of only letters and digits.   * Required if:
    backup_quota int        // A positive integer value (i.e. 1 and up).   * Required if:
    bandwidth_quota int     // A non-negative integer value (i.e. 0 and up).   * Optional
    config_id int           // A positive integer value (i.e. 1 and up).   * Required
    domain string           // A fully-qualified domain name (i.e. liquidweb.com, www.liquidweb.com, etc)  * Required
    image_id int            // A positive integer value (i.e. 1 and up).   
    ip_count int            // A positive integer value (i.e. 1 and up).   * Optional
    ms_sql string           // A string.   
    password string         // A password of at least 7 characters and up to 30 characters in length, containing no spaces. Must contain 3 of the 4 following classes: lowercase, uppercase, numbers and punctuation.  * Required
    public_ssh_key string   // A public ssh key, e.g. the contents of an id_rsa.pub or id_dsa.pub file.    
    template string         // A single word, consisting of only letters, digits, and underscores. * Required if:
    zone int                // A positive integer value (i.e. 1 and up).   
}


type StormCreateResponse struct {
    server StormServer
}


type StormDestroyRequest struct {
    uniq_id string          // A six-character identifier, containing only capital letters and digits. * Required
}

type StormDestroyResponse struct {
    destroyed string        // A six-character identifier, containing only capital letters and digits.
}


type StormDetailsRequest struct {
    uniq_id string          // A six-character identifier, containing only capital letters and digits. * Required
}

type StormDetailResponse struct {
    server StormServer
}

type StormListRequest struct {
    page_num int    // A positive integer value (i.e. 1 and up).
    page_size int   // A positive integer value (i.e. 1 and up).
}

type StormList struct {
    item_count int // A non-negative integer value (i.e. 0 and up).
    item_total int // A non-negative integer value (i.e. 0 and up).
    page_num int // A positive integer value (i.e. 1 and up).
    page_size int // A positive integer value (i.e. 1 and up).
    page_total int // A positive integer value (i.e. 1 and up).
    items map[string]StormServer // An array of associative arrays containing
}

type StormServer struct {
    accnt int                       // A valid account number, up to 6 digits in length.   
    active bool                     // A boolean value (0 or 1).   
    backup_enabled bool             // A boolean value (0 or 1).   
    backup_plan string              // A single word, consisting of only letters, digits, and underscores. * Required if:
    backup_quota int                // A non-negative integer value (i.e. 0 and up).   * Required if:
    backup_size float               // A floating-point value. * Optional
    bandwidth_quota int             // A non-negative integer value (i.e. 0 and up).   
    config_description string       // A string.   
    config_id int                   // A positive integer value (i.e. 1 and up).   
    create_date string              // A valid date and time in YYYY-MM-DD HH:MM:SS format.    
    domain string                   // A fully-qualified domain name (i.e. liquidweb.com, www.liquidweb.com, etc)  
    ip string                       // An IPv4 IP address, in quad-dotted decimal notation (i.e. 127.0.0.1)    
    ip_count int                    // A positive integer value (i.e. 1 and up).   
    manage_level string             // A single word, consisting of only letters and digits.   
    template string                 // A single word, consisting of only letters, digits, and underscores. 
    template_description string     // A string.   
    type string                     // A valid product code, at least 3 chars long, including at least one letter. 
    uniq_id string                  // A six-character identifier, containing only capital letters and digits. 
    zone StormZone                  // An associative array containing the values: * Optional
}


type StormZone struct {
    id int
    name string
    region StormRegion
}

type StormRegion struct {
    id int
    name strinf
}


type StormAPI struct {
    token string
    expires string
}


func NewStormAPI(timeout int) *StormAPI {

    // authenticate

    return &StormAPI{token: token, expires: expires}
}


func (s *StormAPI) clone() StormServer {

}

func (s *StormAPI) create() StormServer {

}

func (s *StormAPI) details() StormServer {

}

func (s *StormAPI) list() StormList {

}


