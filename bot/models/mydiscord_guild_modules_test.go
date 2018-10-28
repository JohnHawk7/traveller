// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMydiscordGuildModules(t *testing.T) {
	t.Parallel()

	query := MydiscordGuildModules()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMydiscordGuildModulesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMydiscordGuildModulesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := MydiscordGuildModules().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMydiscordGuildModulesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MydiscordGuildModuleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMydiscordGuildModulesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MydiscordGuildModuleExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if MydiscordGuildModule exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MydiscordGuildModuleExists to return true, but got false.")
	}
}

func testMydiscordGuildModulesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	mydiscordGuildModuleFound, err := FindMydiscordGuildModule(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if mydiscordGuildModuleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMydiscordGuildModulesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = MydiscordGuildModules().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMydiscordGuildModulesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := MydiscordGuildModules().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMydiscordGuildModulesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	mydiscordGuildModuleOne := &MydiscordGuildModule{}
	mydiscordGuildModuleTwo := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, mydiscordGuildModuleOne, mydiscordGuildModuleDBTypes, false, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}
	if err = randomize.Struct(seed, mydiscordGuildModuleTwo, mydiscordGuildModuleDBTypes, false, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mydiscordGuildModuleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mydiscordGuildModuleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MydiscordGuildModules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMydiscordGuildModulesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	mydiscordGuildModuleOne := &MydiscordGuildModule{}
	mydiscordGuildModuleTwo := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, mydiscordGuildModuleOne, mydiscordGuildModuleDBTypes, false, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}
	if err = randomize.Struct(seed, mydiscordGuildModuleTwo, mydiscordGuildModuleDBTypes, false, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = mydiscordGuildModuleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = mydiscordGuildModuleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func mydiscordGuildModuleBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuildModule) error {
	*o = MydiscordGuildModule{}
	return nil
}

func mydiscordGuildModuleAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuildModule) error {
	*o = MydiscordGuildModule{}
	return nil
}

func mydiscordGuildModuleAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuildModule) error {
	*o = MydiscordGuildModule{}
	return nil
}

func mydiscordGuildModuleBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuildModule) error {
	*o = MydiscordGuildModule{}
	return nil
}

func mydiscordGuildModuleAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuildModule) error {
	*o = MydiscordGuildModule{}
	return nil
}

func mydiscordGuildModuleBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuildModule) error {
	*o = MydiscordGuildModule{}
	return nil
}

func mydiscordGuildModuleAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuildModule) error {
	*o = MydiscordGuildModule{}
	return nil
}

func mydiscordGuildModuleBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuildModule) error {
	*o = MydiscordGuildModule{}
	return nil
}

func mydiscordGuildModuleAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *MydiscordGuildModule) error {
	*o = MydiscordGuildModule{}
	return nil
}

func testMydiscordGuildModulesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &MydiscordGuildModule{}
	o := &MydiscordGuildModule{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule object: %s", err)
	}

	AddMydiscordGuildModuleHook(boil.BeforeInsertHook, mydiscordGuildModuleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildModuleBeforeInsertHooks = []MydiscordGuildModuleHook{}

	AddMydiscordGuildModuleHook(boil.AfterInsertHook, mydiscordGuildModuleAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildModuleAfterInsertHooks = []MydiscordGuildModuleHook{}

	AddMydiscordGuildModuleHook(boil.AfterSelectHook, mydiscordGuildModuleAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildModuleAfterSelectHooks = []MydiscordGuildModuleHook{}

	AddMydiscordGuildModuleHook(boil.BeforeUpdateHook, mydiscordGuildModuleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildModuleBeforeUpdateHooks = []MydiscordGuildModuleHook{}

	AddMydiscordGuildModuleHook(boil.AfterUpdateHook, mydiscordGuildModuleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildModuleAfterUpdateHooks = []MydiscordGuildModuleHook{}

	AddMydiscordGuildModuleHook(boil.BeforeDeleteHook, mydiscordGuildModuleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildModuleBeforeDeleteHooks = []MydiscordGuildModuleHook{}

	AddMydiscordGuildModuleHook(boil.AfterDeleteHook, mydiscordGuildModuleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildModuleAfterDeleteHooks = []MydiscordGuildModuleHook{}

	AddMydiscordGuildModuleHook(boil.BeforeUpsertHook, mydiscordGuildModuleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildModuleBeforeUpsertHooks = []MydiscordGuildModuleHook{}

	AddMydiscordGuildModuleHook(boil.AfterUpsertHook, mydiscordGuildModuleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	mydiscordGuildModuleAfterUpsertHooks = []MydiscordGuildModuleHook{}
}

func testMydiscordGuildModulesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMydiscordGuildModulesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(mydiscordGuildModuleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMydiscordGuildModuleToOneMydiscordGuildUsingGuild(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local MydiscordGuildModule
	var foreign MydiscordGuild

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, mydiscordGuildModuleDBTypes, false, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, mydiscordGuildDBTypes, false, mydiscordGuildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuild struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.GuildID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Guild().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := MydiscordGuildModuleSlice{&local}
	if err = local.L.LoadGuild(ctx, tx, false, (*[]*MydiscordGuildModule)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Guild == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Guild = nil
	if err = local.L.LoadGuild(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Guild == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMydiscordGuildModuleToOneMydiscordModuleUsingModule(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local MydiscordGuildModule
	var foreign MydiscordModule

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, mydiscordGuildModuleDBTypes, false, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, mydiscordModuleDBTypes, false, mydiscordModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordModule struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ModuleID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Module().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := MydiscordGuildModuleSlice{&local}
	if err = local.L.LoadModule(ctx, tx, false, (*[]*MydiscordGuildModule)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Module == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Module = nil
	if err = local.L.LoadModule(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Module == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMydiscordGuildModuleToOneSetOpMydiscordGuildUsingGuild(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MydiscordGuildModule
	var b, c MydiscordGuild

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mydiscordGuildModuleDBTypes, false, strmangle.SetComplement(mydiscordGuildModulePrimaryKeyColumns, mydiscordGuildModuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mydiscordGuildDBTypes, false, strmangle.SetComplement(mydiscordGuildPrimaryKeyColumns, mydiscordGuildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mydiscordGuildDBTypes, false, strmangle.SetComplement(mydiscordGuildPrimaryKeyColumns, mydiscordGuildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MydiscordGuild{&b, &c} {
		err = a.SetGuild(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Guild != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.GuildMydiscordGuildModules[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GuildID != x.ID {
			t.Error("foreign key was wrong value", a.GuildID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GuildID))
		reflect.Indirect(reflect.ValueOf(&a.GuildID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GuildID != x.ID {
			t.Error("foreign key was wrong value", a.GuildID, x.ID)
		}
	}
}
func testMydiscordGuildModuleToOneSetOpMydiscordModuleUsingModule(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a MydiscordGuildModule
	var b, c MydiscordModule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, mydiscordGuildModuleDBTypes, false, strmangle.SetComplement(mydiscordGuildModulePrimaryKeyColumns, mydiscordGuildModuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, mydiscordModuleDBTypes, false, strmangle.SetComplement(mydiscordModulePrimaryKeyColumns, mydiscordModuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, mydiscordModuleDBTypes, false, strmangle.SetComplement(mydiscordModulePrimaryKeyColumns, mydiscordModuleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*MydiscordModule{&b, &c} {
		err = a.SetModule(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Module != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ModuleMydiscordGuildModules[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ModuleID != x.ID {
			t.Error("foreign key was wrong value", a.ModuleID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ModuleID))
		reflect.Indirect(reflect.ValueOf(&a.ModuleID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ModuleID != x.ID {
			t.Error("foreign key was wrong value", a.ModuleID, x.ID)
		}
	}
}

func testMydiscordGuildModulesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMydiscordGuildModulesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MydiscordGuildModuleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMydiscordGuildModulesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := MydiscordGuildModules().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	mydiscordGuildModuleDBTypes = map[string]string{`GuildID`: `integer`, `ID`: `integer`, `ModuleID`: `integer`}
	_                           = bytes.MinRead
)

func testMydiscordGuildModulesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(mydiscordGuildModulePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(mydiscordGuildModuleColumns) == len(mydiscordGuildModulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMydiscordGuildModulesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(mydiscordGuildModuleColumns) == len(mydiscordGuildModulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &MydiscordGuildModule{}
	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModuleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, mydiscordGuildModuleDBTypes, true, mydiscordGuildModulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(mydiscordGuildModuleColumns, mydiscordGuildModulePrimaryKeyColumns) {
		fields = mydiscordGuildModuleColumns
	} else {
		fields = strmangle.SetComplement(
			mydiscordGuildModuleColumns,
			mydiscordGuildModulePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MydiscordGuildModuleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMydiscordGuildModulesUpsert(t *testing.T) {
	t.Parallel()

	if len(mydiscordGuildModuleColumns) == len(mydiscordGuildModulePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := MydiscordGuildModule{}
	if err = randomize.Struct(seed, &o, mydiscordGuildModuleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MydiscordGuildModule: %s", err)
	}

	count, err := MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, mydiscordGuildModuleDBTypes, false, mydiscordGuildModulePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize MydiscordGuildModule struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert MydiscordGuildModule: %s", err)
	}

	count, err = MydiscordGuildModules().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
