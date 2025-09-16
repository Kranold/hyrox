package helperfunctions

import "database/sql"

// ToNullString converts a string to sql.NullString
func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

// ToNullInt32 converts an int to sql.NullInt32
func ToNullInt32(i int32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{Int32: 0, Valid: false}
	}
	return sql.NullInt32{Int32: int32(i), Valid: true}
}

func ToNullInt32FromInt(i int) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{Int32: 0, Valid: false}
	}
	return sql.NullInt32{Int32: int32(i), Valid: true}
}

func ToNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}
	return sql.NullInt64{Int64: int64(i), Valid: true}
}

func ToNullFloat64(f float64) sql.NullFloat64 {
	if f == 0 {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
	return sql.NullFloat64{Float64: f, Valid: true}
}

func ToNullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}
