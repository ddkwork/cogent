// Code generated by "core generate"; DO NOT EDIT.

package mail

import (
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

// AppType is the [types.Type] for [App]
var AppType = types.AddType(&types.Type{Name: "cogentcore.org/cogent/mail.App", IDName: "app", Doc: "App is an email client app.", Methods: []types.Method{{Name: "MoveMessage", Doc: "MoveMessage moves the current message to the given mailbox.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"mailbox"}, Returns: []string{"error"}}, {Name: "Compose", Doc: "Compose pulls up a dialog to send a new message", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SendMessage", Doc: "SendMessage sends the current message", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Returns: []string{"error"}}}, Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "AuthToken", Doc: "AuthToken contains the [oauth2.Token] for each account."}, {Name: "AuthClient", Doc: "AuthClient contains the [sasl.Client] authentication for sending messages for each account."}, {Name: "IMAPClient", Doc: "IMAPCLient contains the imap clients for each account."}, {Name: "ComposeMessage", Doc: "ComposeMessage is the current message we are editing"}, {Name: "Cache", Doc: "Cache contains the cache data, keyed by account and then mailbox."}, {Name: "ReadMessage", Doc: "ReadMessage is the current message we are reading"}, {Name: "CurrentEmail", Doc: "The current email account"}, {Name: "CurrentMailbox", Doc: "The current mailbox"}}, Instance: &App{}})

// NewApp returns a new [App] with the given optional parent:
// App is an email client app.
func NewApp(parent ...tree.Node) *App { return tree.New[*App](parent...) }

// NodeType returns the [*types.Type] of [App]
func (t *App) NodeType() *types.Type { return AppType }

// New returns a new [*App] value
func (t *App) New() tree.Node { return &App{} }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/cogent/mail.SettingsData", IDName: "settings-data", Doc: "SettingsData is the data type for the global Cogent Mail settings.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Embeds: []types.Field{{Name: "SettingsBase"}}, Fields: []types.Field{{Name: "Accounts", Doc: "Accounts are the email accounts the user is signed into."}}})

// AddressTextFieldType is the [types.Type] for [AddressTextField]
var AddressTextFieldType = types.AddType(&types.Type{Name: "cogentcore.org/cogent/mail.AddressTextField", IDName: "address-text-field", Doc: "AddressTextField represents a [mail.Address] with a [core.TextField].", Embeds: []types.Field{{Name: "TextField"}}, Instance: &AddressTextField{}})

// NewAddressTextField returns a new [AddressTextField] with the given optional parent:
// AddressTextField represents a [mail.Address] with a [core.TextField].
func NewAddressTextField(parent ...tree.Node) *AddressTextField {
	return tree.New[*AddressTextField](parent...)
}

// NodeType returns the [*types.Type] of [AddressTextField]
func (t *AddressTextField) NodeType() *types.Type { return AddressTextFieldType }

// New returns a new [*AddressTextField] value
func (t *AddressTextField) New() tree.Node { return &AddressTextField{} }
