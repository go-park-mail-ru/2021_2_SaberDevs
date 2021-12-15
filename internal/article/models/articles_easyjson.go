// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels(in *jlexer.Lexer, out *Preview) {
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
		case "id":
			out.Id = string(in.String())
		case "datetime":
			out.DateTime = string(in.String())
		case "previewUrl":
			out.PreviewUrl = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Tags = append(out.Tags, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "title":
			out.Title = string(in.String())
		case "category":
			out.Category = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "author":
			easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels1(in, &out.Author)
		case "commentsUrl":
			out.CommentsUrl = string(in.String())
		case "comments":
			out.Comments = uint(in.Uint())
		case "likes":
			out.Likes = int64(in.Int64())
		case "liked":
			out.Liked = int64(in.Int64())
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
func easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels(out *jwriter.Writer, in Preview) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"datetime\":"
		out.RawString(prefix)
		out.String(string(in.DateTime))
	}
	{
		const prefix string = ",\"previewUrl\":"
		out.RawString(prefix)
		out.String(string(in.PreviewUrl))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Tags {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels1(out, in.Author)
	}
	{
		const prefix string = ",\"commentsUrl\":"
		out.RawString(prefix)
		out.String(string(in.CommentsUrl))
	}
	{
		const prefix string = ",\"comments\":"
		out.RawString(prefix)
		out.Uint(uint(in.Comments))
	}
	{
		const prefix string = ",\"likes\":"
		out.RawString(prefix)
		out.Int64(int64(in.Likes))
	}
	{
		const prefix string = ",\"liked\":"
		out.RawString(prefix)
		out.Int64(int64(in.Liked))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Preview) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Preview) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Preview) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Preview) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels(l, v)
}
func easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels1(in *jlexer.Lexer, out *Author) {
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
		case "login":
			out.Login = string(in.String())
		case "firstName":
			out.Name = string(in.String())
		case "lastName":
			out.Surname = string(in.String())
		case "avatarUrl":
			out.AvatarUrl = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "score":
			out.Score = int(in.Int())
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
func easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels1(out *jwriter.Writer, in Author) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"firstName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"lastName\":"
		out.RawString(prefix)
		out.String(string(in.Surname))
	}
	{
		const prefix string = ",\"avatarUrl\":"
		out.RawString(prefix)
		out.String(string(in.AvatarUrl))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"score\":"
		out.RawString(prefix)
		out.Int(int(in.Score))
	}
	out.RawByte('}')
}
func easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels2(in *jlexer.Lexer, out *GenericResponse) {
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
		case "status":
			out.Status = uint(in.Uint())
		case "data":
			out.Data = string(in.String())
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
func easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels2(out *jwriter.Writer, in GenericResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Status))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		out.String(string(in.Data))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GenericResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GenericResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GenericResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GenericResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels2(l, v)
}
func easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels3(in *jlexer.Lexer, out *FullArticle) {
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
		case "id":
			out.Id = string(in.String())
		case "datetime":
			out.DateTime = string(in.String())
		case "previewUrl":
			out.PreviewUrl = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Tags = append(out.Tags, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "title":
			out.Title = string(in.String())
		case "category":
			out.Category = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "author":
			easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels1(in, &out.Author)
		case "commentsUrl":
			out.CommentsUrl = string(in.String())
		case "comments":
			out.Comments = uint(in.Uint())
		case "likes":
			out.Likes = int64(in.Int64())
		case "liked":
			out.Liked = int64(in.Int64())
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
func easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels3(out *jwriter.Writer, in FullArticle) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"datetime\":"
		out.RawString(prefix)
		out.String(string(in.DateTime))
	}
	{
		const prefix string = ",\"previewUrl\":"
		out.RawString(prefix)
		out.String(string(in.PreviewUrl))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Tags {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix)
		easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels1(out, in.Author)
	}
	{
		const prefix string = ",\"commentsUrl\":"
		out.RawString(prefix)
		out.String(string(in.CommentsUrl))
	}
	{
		const prefix string = ",\"comments\":"
		out.RawString(prefix)
		out.Uint(uint(in.Comments))
	}
	{
		const prefix string = ",\"likes\":"
		out.RawString(prefix)
		out.Int64(int64(in.Likes))
	}
	{
		const prefix string = ",\"liked\":"
		out.RawString(prefix)
		out.Int64(int64(in.Liked))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FullArticle) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FullArticle) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FullArticle) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FullArticle) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels3(l, v)
}
func easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels4(in *jlexer.Lexer, out *ChunkResponse) {
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
		case "status":
			out.Status = uint(in.Uint())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.ChunkData = nil
			} else {
				in.Delim('[')
				if out.ChunkData == nil {
					if !in.IsDelim(']') {
						out.ChunkData = make([]Preview, 0, 0)
					} else {
						out.ChunkData = []Preview{}
					}
				} else {
					out.ChunkData = (out.ChunkData)[:0]
				}
				for !in.IsDelim(']') {
					var v7 Preview
					(v7).UnmarshalEasyJSON(in)
					out.ChunkData = append(out.ChunkData, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels4(out *jwriter.Writer, in ChunkResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Status))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		if in.ChunkData == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.ChunkData {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ChunkResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ChunkResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ChunkResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ChunkResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels4(l, v)
}
func easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels5(in *jlexer.Lexer, out *AuthorsChunks) {
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
		case "status":
			out.Status = uint(in.Uint())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.ChunkData = nil
			} else {
				in.Delim('[')
				if out.ChunkData == nil {
					if !in.IsDelim(']') {
						out.ChunkData = make([]Author, 0, 0)
					} else {
						out.ChunkData = []Author{}
					}
				} else {
					out.ChunkData = (out.ChunkData)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Author
					easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels1(in, &v10)
					out.ChunkData = append(out.ChunkData, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels5(out *jwriter.Writer, in AuthorsChunks) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Status))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		if in.ChunkData == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.ChunkData {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels1(out, v12)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthorsChunks) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthorsChunks) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthorsChunks) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthorsChunks) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels5(l, v)
}
func easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels6(in *jlexer.Lexer, out *ArticleUpdate) {
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
		case "id":
			out.Id = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "category":
			out.Category = string(in.String())
		case "img":
			out.Img = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.Tags = append(out.Tags, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels6(out *jwriter.Writer, in ArticleUpdate) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"img\":"
		out.RawString(prefix)
		out.String(string(in.Img))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Tags {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ArticleUpdate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ArticleUpdate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ArticleUpdate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ArticleUpdate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels6(l, v)
}
func easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels7(in *jlexer.Lexer, out *ArticleResponse) {
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
		case "status":
			out.Status = uint(in.Uint())
		case "data":
			(out.Data).UnmarshalEasyJSON(in)
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
func easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels7(out *jwriter.Writer, in ArticleResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Status))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		(in.Data).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ArticleResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ArticleResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ArticleResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ArticleResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels7(l, v)
}
func easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels8(in *jlexer.Lexer, out *ArticleCreate) {
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
		case "title":
			out.Title = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "category":
			out.Category = string(in.String())
		case "img":
			out.Img = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]string, 0, 4)
					} else {
						out.Tags = []string{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.Tags = append(out.Tags, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels8(out *jwriter.Writer, in ArticleCreate) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"img\":"
		out.RawString(prefix)
		out.String(string(in.Img))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Tags {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ArticleCreate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ArticleCreate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEaa7cc45EncodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ArticleCreate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ArticleCreate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEaa7cc45DecodeGithubComGoParkMailRu20212SaberDevsInternalArticleModels8(l, v)
}
