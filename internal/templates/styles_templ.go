// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Styles() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\r\n        body {\r\n            background-color: #000000;\r\n            color: #ffffff;\r\n            font-family: Arial, sans-serif;\r\n            margin: 0;\r\n            display: flex;\r\n            flex-direction: column;\r\n            align-items: center;\r\n            min-height: 100vh;\r\n        }\r\n\r\n        main {\r\n            max-width: 1200px;\r\n            padding: 20px;\r\n            display: flex;\r\n            flex-direction: column;\r\n            align-items: center;\r\n        }\r\n\r\n        .title {\r\n            margin-top: 40px;\r\n            text-align: center;\r\n            font-size: 50px;\r\n            margin-bottom: 20px;\r\n            color: #ff6600;\r\n        }\r\n\r\n        h1 {\r\n            font-size: 32px;\r\n        }\r\n\r\n        section {\r\n            width: 100%;\r\n            margin-bottom: 40px;\r\n        }\r\n\r\n        form {\r\n            display: flex;\r\n            flex-direction: column;\r\n            align-items: center;\r\n            width: 100%;\r\n        }\r\n\r\n        input[type=\"file\"] {\r\n            padding: 30px;\r\n            margin-bottom: 20px;\r\n            width: 100%;\r\n            border-radius: 10px;\r\n            background-color: #1a1a1a;\r\n            color: #ff6600;\r\n            font-size: 30px;\r\n            display: block;\r\n            font-family: monospace;\r\n            font-weight: bold;\r\n            text-align: center;\r\n        }\r\n        input[type=\"file\"]:hover {\r\n            background-color: #333;\r\n            cursor: pointer;\r\n        }\r\n\r\n        button {\r\n            background-color: #ff6600;\r\n            color: #ffffff;\r\n            border: none;\r\n            padding: 10px 20px;\r\n            margin-bottom: 20px;\r\n            cursor: pointer;\r\n            border-radius: 4px;\r\n            font-weight: bold;\r\n        }\r\n        button:hover {\r\n            background-color: #b44800;\r\n        }\r\n\r\n        #output-container {\r\n            width: 100%;\r\n        }\r\n\r\n        #output {\r\n            width: 100%;\r\n            padding: 20px 0;\r\n            background-color: #1a1a1a;\r\n            border-radius: 8px;\r\n            white-space: pre;\r\n            word-wrap: break-word;\r\n            border: 1px solid #333;\r\n            font-family: monospace;\r\n            font-size: 2.5px;\r\n            text-align: center;\r\n            display: flex;\r\n            flex-direction: column;\r\n            gap: 20px;\r\n            align-items: center;\r\n            justify-content: center;\r\n            overflow-x: auto;\r\n        }\r\n\r\n        footer {\r\n            text-align: center;\r\n            padding: 20px;\r\n            font-size: 16px;\r\n            color: #888888;\r\n            margin-top: auto;\r\n            display: flex;\r\n            flex-direction: row;\r\n            justify-content: space-between;\r\n            width: 90%;\r\n        }\r\n\r\n        ul {\r\n            list-style-type: none;\r\n            display: flex;\r\n            justify-content: center;\r\n            align-items: center;\r\n            flex-direction: row;\r\n            gap: 10px;\r\n        }\r\n\r\n        a {\r\n            color: #ff6600;\r\n            text-decoration: none;\r\n        }\r\n\r\n        footer a:hover {\r\n            text-decoration: underline;\r\n        }\r\n\r\n        @media (max-width: 768px) {\r\n            main {\r\n                padding: 10px;\r\n            }\r\n\r\n            #output {\r\n                font-size: 1.5px;\r\n            }\r\n\r\n            #output-container {\r\n                overflow-x: auto;\r\n                -webkit-overflow-scrolling: touch; /* For smooth scrolling on iOS */\r\n            }\r\n            \r\n            footer {\r\n                padding: 10px;\r\n            }\r\n        }\r\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
