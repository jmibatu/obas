package io

import (
	"errors"
	"obas/src/api"
	domain "obas/src/domain/demographics"
)

const raceUrl = api.BASE_URL + "/demographics"

type Races domain.Race

func GetRaces() ([]domain.Race, error) {
	entites := []domain.Race{}
	resp, _ := api.Rest().Get(raceUrl + "/all")

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetRace(id string) (domain.Race, error) {
	entity := domain.Race{}
	resp, _ := api.Rest().Get(raceUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateRace(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(raceUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateRace(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(raceUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteRace(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(raceUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
