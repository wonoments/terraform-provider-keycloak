package provider

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tazjin/terraform-provider-keycloak/keycloak"
)

func resourceRealm() *schema.Resource {
	return &schema.Resource{
		// API methods
		Read:   schema.ReadFunc(resourceRealmRead),
		Create: schema.CreateFunc(resourceRealmCreate),
		Update: schema.UpdateFunc(resourceRealmUpdate),
		Delete: schema.DeleteFunc(resourceRealmDelete),

		// Realms are importable by ID
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"realm": {
				Description: "Realm name and ID",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"ssl_required": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "EXTERNAL",
				ValidateFunc: func(v interface{}, _ string) (w []string, err []error) {
					switch v.(string) {
					case
						"ALL",
						"EXTERNAL",
						"NONE":
						return
					}
					err = []error{
						fmt.Errorf("Invalid value for ssl_required. Valid are ALL, EXTERNAL or NONE"),
					}
					return
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"supported_locales": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"default_roles": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"smtp_server": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"internationalization_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"registration_allowed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"registration_email_as_username": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"remember_me": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"verify_email": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"reset_password_allowed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"edit_username_allowed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"brute_force_protected": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"access_token_lifespan": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"access_token_lifespan_for_implicit_flow": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sso_session_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sso_session_max_lifespan": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"offline_session_idle_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"access_code_lifespan": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"access_code_lifespan_user_action": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"access_code_lifespan_login": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_failure_wait_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minimum_quick_login_wait_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"wait_increment_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"quick_login_check_milli_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_delta_time_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"failure_factor": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceRealmRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*keycloak.KeycloakClient)

	r, err := c.GetRealm(d.Id())
	if err != nil {
		return err
	}

	realmToResourceData(r, d)
	return nil
}

func resourceRealmCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*keycloak.KeycloakClient)
	r := resourceDataToRealm(d)

	created, err := c.CreateRealm(r)
	if err != nil {
		return err
	}

	d.SetId(created.Id)

	return resourceRealmRead(d, m)
}

func resourceRealmUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*keycloak.KeycloakClient)
	r := resourceDataToRealm(d)
	return c.UpdateRealm(r)
}

func resourceRealmDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*keycloak.KeycloakClient)
	return c.DeleteRealm(d.Id())
}

// Type/struct conversion boilerplate (thanks, Go)

func resourceDataToRealm(d *schema.ResourceData) *keycloak.Realm {
	r := keycloak.Realm{
		Realm:   d.Get("realm").(string),
		Enabled: d.Get("enabled").(bool),

		SslRequired:      d.Get("ssl_required").(string),
		DisplayName:      d.Get("display_name").(string),
		SupportedLocales: getStringSlice(d, "supported_locales"),
		DefaultRoles:     getStringSlice(d, "default_roles"),

		InternationalizationEnabled: getOptionalBool(d, "internationalization_enabled"),
		RegistrationAllowed:         getOptionalBool(d, "registration_allowed"),
		RegistrationEmailAsUsername: getOptionalBool(d, "registration_email_as_username"),
		RememberMe:                  getOptionalBool(d, "remember_me"),
		VerifyEmail:                 getOptionalBool(d, "verify_email"),
		ResetPasswordAllowed:        getOptionalBool(d, "reset_password_allowed"),
		EditUsernameAllowed:         getOptionalBool(d, "edit_username_allowed"),
		BruteForceProtected:         getOptionalBool(d, "brute_force_protected"),

		AccessTokenLifespan:                getOptionalInt(d, "access_token_lifespan"),
		AccessTokenLifespanForImplicitFlow: getOptionalInt(d, "access_token_lifespan_for_implicit_flow"),
		SsoSessionIdleTimeout:              getOptionalInt(d, "sso_session_idle_timeout"),
		SsoSessionMaxLifespan:              getOptionalInt(d, "sso_session_max_lifespan"),
		OfflineSessionIdleTimeout:          getOptionalInt(d, "offline_session_idle_timeout"),
		AccessCodeLifespan:                 getOptionalInt(d, "access_code_lifespan"),
		AccessCodeLifespanUserAction:       getOptionalInt(d, "access_code_lifespan_user_action"),
		AccessCodeLifespanLogin:            getOptionalInt(d, "access_code_lifespan_login"),
		MaxFailureWaitSeconds:              getOptionalInt(d, "max_failure_wait_seconds"),
		MinimumQuickLoginWaitSeconds:       getOptionalInt(d, "minimum_quick_login_wait_seconds"),
		WaitIncrementSeconds:               getOptionalInt(d, "wait_increment_seconds"),
		QuickLoginCheckMilliSeconds:        getOptionalInt(d, "quick_login_check_milli_seconds"),
		MaxDeltaTimeSeconds:                getOptionalInt(d, "max_delta_time_seconds"),
		FailureFactor:                      getOptionalInt(d, "failure_factor"),
	}

	if !d.IsNewResource() {
		r.Id = d.Id()
	}

	if smtpMap, present := d.GetOk("smtp_server"); present {
		smtp := smtpMap.(keycloak.SmtpServer)
		r.SmtpServer = &smtp
	}

	return &r
}

func realmToResourceData(r *keycloak.Realm, d *schema.ResourceData) {
	d.SetId(r.Id)
	d.Set("realm", r.Realm)
	d.Set("enabled", r.Enabled)

	d.Set("ssl_required", r.SslRequired)
	d.Set("display_name", r.DisplayName)
	d.Set("supported_locales", r.SupportedLocales)
	d.Set("default_roles", r.DefaultRoles)

	if r.SmtpServer != nil {
		d.Set("smtp_server", *r.SmtpServer)
	}

	setOptionalBool(d, "internationalization_enabled", r.InternationalizationEnabled)
	setOptionalBool(d, "registration_allowed", r.RegistrationAllowed)
	setOptionalBool(d, "registration_email_as_username", r.RegistrationEmailAsUsername)
	setOptionalBool(d, "remember_me", r.RememberMe)
	setOptionalBool(d, "verify_email", r.VerifyEmail)
	setOptionalBool(d, "reset_password_allowed", r.ResetPasswordAllowed)
	setOptionalBool(d, "edit_username_allowed", r.EditUsernameAllowed)
	setOptionalBool(d, "brute_force_protected", r.BruteForceProtected)

	setOptionalInt(d, "access_token_lifespan", r.AccessTokenLifespan)
	setOptionalInt(d, "access_token_lifespan_for_implicit_flow", r.AccessTokenLifespanForImplicitFlow)
	setOptionalInt(d, "sso_session_idle_timeout", r.SsoSessionIdleTimeout)
	setOptionalInt(d, "sso_session_max_lifespan", r.SsoSessionMaxLifespan)
	setOptionalInt(d, "offline_session_idle_timeout", r.OfflineSessionIdleTimeout)
	setOptionalInt(d, "access_code_lifespan", r.AccessCodeLifespan)
	setOptionalInt(d, "access_code_lifespan_user_action", r.AccessCodeLifespanUserAction)
	setOptionalInt(d, "access_code_lifespan_login", r.AccessCodeLifespanLogin)
	setOptionalInt(d, "max_failure_wait_seconds", r.MaxFailureWaitSeconds)
	setOptionalInt(d, "minimum_quick_login_wait_seconds", r.MinimumQuickLoginWaitSeconds)
	setOptionalInt(d, "wait_increment_seconds", r.WaitIncrementSeconds)
	setOptionalInt(d, "quick_login_check_milli_seconds", r.QuickLoginCheckMilliSeconds)
	setOptionalInt(d, "max_delta_time_seconds", r.MaxDeltaTimeSeconds)
	setOptionalInt(d, "failure_factor", r.FailureFactor)
}
