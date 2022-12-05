// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package chromedpundetected

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6601e8cdDecodeGithubComDavincibleChromedpUndetected(in *jlexer.Lexer, out *userAgent) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "userAgent":
			out.UserAgent = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6601e8cdEncodeGithubComDavincibleChromedpUndetected(out *jwriter.Writer, in userAgent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userAgent\":"
		out.RawString(prefix[1:])
		out.String(string(in.UserAgent))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v userAgent) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6601e8cdEncodeGithubComDavincibleChromedpUndetected(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v userAgent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6601e8cdEncodeGithubComDavincibleChromedpUndetected(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *userAgent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6601e8cdDecodeGithubComDavincibleChromedpUndetected(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *userAgent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6601e8cdDecodeGithubComDavincibleChromedpUndetected(l, v)
}