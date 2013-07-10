package grid

import(
	"fmt"
	"github.com/timcubb/juno/system"
    "github.com/timcubb/juno/deploy"
	)

type Machine struct {
    Name string   //
    Id string     // storm servers ID for server
    Type int      // 8/16/32/64
    Hostname string
    Status int
}

type MachineRequest struct {
    config_id int       //   A positive integer value (i.e. 1 and up).   
    domain string       //A fully-qualified domain name (i.e. liquidweb.com, www.liquidweb.com, etc)  * Required
    ip_count int        //A positive integer value (i.e. 1 and up).   
    password string     // A password of at least 7 characters and up to 30 characters in length, containing no spaces. Must contain 3 of the 4 following classes: lowercase, uppercase, numbers and punctuation.  * Required
    uniq_id string      // A six-character identifier, containing only capital letters and digits. * Required
    zone int            // A non-negative integer value (i.e. 0 and up).
}


type MachineResponse struct {
    accnt int                   // A valid account number, up to 6 digits in length.   
    active bool                 //A boolean value (0 or 1).   
    backup_enabled bool         // A boolean value (0 or 1).   
    backup_plan string          // A single word, consisting of only letters, digits, and underscores. * Required if:
    backup_enabled bool         // is set to '1'
    backup_quota int            // A non-negative integer value (i.e. 0 and up).   * Required if:
    backup_enabled bool         // is set to '1'
    backup_size float           // A floating-point value. * Optional
    bandwidth_quota int         // A non-negative integer value (i.e. 0 and up).   
    config_description string   // A string.   
    config_id int               // A positive integer value (i.e. 1 and up).   
    create_date string          // A valid date and time in YYYY-MM-DD HH:MM:SS format.    
    domain string               // A fully-qualified domain name (i.e. liquidweb.com, www.liquidweb.com, etc)  
    ip string                   // An IPv4 IP address, in quad-dotted decimal notation (i.e. 127.0.0.1)    
    ip_count int                // A positive integer value (i.e. 1 and up).   
    manage_level string         // A single word, consisting of only letters and digits.   
    template string             // A single word, consisting of only letters, digits, and underscores. 
    template_description string // A string.   
    type string                 // A valid product code, at least 3 chars long, including at least one letter. 
    uniq_id string              // A six-character identifier, containing only capital letters and digits. 
    zone string                 // An associative array containing the values: * Optional
}