package handlers

//PutTrackingInfo atualiza as informações de status de um determinado tracking
// func PutTrackingInfo(w http.ResponseWriter, r *http.Request) {
// 	ps := r.Context().Value("params").(httprouter.Params) //Preciso obter os parâmetros a partir do contexto

// 	info := models.TrackingUpdateInfo{}

// 	//Obtém o registro através do id
// 	id, err := strconv.Atoi(ps.ByName("id"))

// 	if err != nil {
// 		utils.SendJSONBadRequest(w, "Informe um Id válido.")
// 		return
// 	}

// 	b, err := utils.ReadBody(r)
// 	if err != nil {
// 		utils.SendJSONBadRequest(w, err.Error())
// 		return
// 	}

// 	err = json.Unmarshal(b, &info)
// 	if err != nil {
// 		utils.SendJSONBadRequest(w, err.Error())
// 		return
// 	}

// 	if err = repo.UpdateTrackingInfo(id, info); err != nil {
// 		switch err {
// 		case errors.ErrIDNotFound:
// 			utils.SendJSONNotFound(w, err.Error())
// 		default:
// 			utils.SendJSONBadRequest(w, err.Error())
// 		}

// 	} else {
// 		utils.SendJSONSuccess(w, "Informações atualizadas com sucesso.")
// 	}
// }
