// Code generated by cdpgen. DO NOT EDIT.

// Package browser implements the Browser domain. The Browser domain defines
// methods and events for browser managing.
package browser

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Browser domain. The Browser domain defines
// methods and events for browser managing.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Browser domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// SetPermission invokes the Browser method. Set permission settings for given
// origin.
func (d *domainClient) SetPermission(ctx context.Context, args *SetPermissionArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Browser.setPermission", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Browser.setPermission", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "SetPermission", Err: err}
	}
	return
}

// GrantPermissions invokes the Browser method. Grant specific permissions to
// the given origin and reject all others.
func (d *domainClient) GrantPermissions(ctx context.Context, args *GrantPermissionsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Browser.grantPermissions", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Browser.grantPermissions", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "GrantPermissions", Err: err}
	}
	return
}

// ResetPermissions invokes the Browser method. Reset all permission
// management for all origins.
func (d *domainClient) ResetPermissions(ctx context.Context, args *ResetPermissionsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Browser.resetPermissions", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Browser.resetPermissions", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "ResetPermissions", Err: err}
	}
	return
}

// Close invokes the Browser method. Close browser gracefully.
func (d *domainClient) Close(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Browser.close", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "Close", Err: err}
	}
	return
}

// Crash invokes the Browser method. Crashes browser on the main thread.
func (d *domainClient) Crash(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Browser.crash", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "Crash", Err: err}
	}
	return
}

// CrashGPUProcess invokes the Browser method. Crashes GPU process.
func (d *domainClient) CrashGPUProcess(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Browser.crashGpuProcess", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "CrashGPUProcess", Err: err}
	}
	return
}

// GetVersion invokes the Browser method. Returns version information.
func (d *domainClient) GetVersion(ctx context.Context) (reply *GetVersionReply, err error) {
	reply = new(GetVersionReply)
	err = rpcc.Invoke(ctx, "Browser.getVersion", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "GetVersion", Err: err}
	}
	return
}

// GetBrowserCommandLine invokes the Browser method. Returns the command line
// switches for the browser process if, and only if --enable-automation is on
// the commandline.
func (d *domainClient) GetBrowserCommandLine(ctx context.Context) (reply *GetBrowserCommandLineReply, err error) {
	reply = new(GetBrowserCommandLineReply)
	err = rpcc.Invoke(ctx, "Browser.getBrowserCommandLine", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "GetBrowserCommandLine", Err: err}
	}
	return
}

// GetHistograms invokes the Browser method. Get Chrome histograms.
func (d *domainClient) GetHistograms(ctx context.Context, args *GetHistogramsArgs) (reply *GetHistogramsReply, err error) {
	reply = new(GetHistogramsReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Browser.getHistograms", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Browser.getHistograms", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "GetHistograms", Err: err}
	}
	return
}

// GetHistogram invokes the Browser method. Get a Chrome histogram by name.
func (d *domainClient) GetHistogram(ctx context.Context, args *GetHistogramArgs) (reply *GetHistogramReply, err error) {
	reply = new(GetHistogramReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Browser.getHistogram", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Browser.getHistogram", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "GetHistogram", Err: err}
	}
	return
}

// GetWindowBounds invokes the Browser method. Get position and size of the
// browser window.
func (d *domainClient) GetWindowBounds(ctx context.Context, args *GetWindowBoundsArgs) (reply *GetWindowBoundsReply, err error) {
	reply = new(GetWindowBoundsReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Browser.getWindowBounds", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Browser.getWindowBounds", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "GetWindowBounds", Err: err}
	}
	return
}

// GetWindowForTarget invokes the Browser method. Get the browser window that
// contains the devtools target.
func (d *domainClient) GetWindowForTarget(ctx context.Context, args *GetWindowForTargetArgs) (reply *GetWindowForTargetReply, err error) {
	reply = new(GetWindowForTargetReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Browser.getWindowForTarget", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Browser.getWindowForTarget", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "GetWindowForTarget", Err: err}
	}
	return
}

// SetWindowBounds invokes the Browser method. Set position and/or size of the
// browser window.
func (d *domainClient) SetWindowBounds(ctx context.Context, args *SetWindowBoundsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Browser.setWindowBounds", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Browser.setWindowBounds", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "SetWindowBounds", Err: err}
	}
	return
}

// SetDockTile invokes the Browser method. Set dock tile details,
// platform-specific.
func (d *domainClient) SetDockTile(ctx context.Context, args *SetDockTileArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Browser.setDockTile", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Browser.setDockTile", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Browser", Op: "SetDockTile", Err: err}
	}
	return
}
