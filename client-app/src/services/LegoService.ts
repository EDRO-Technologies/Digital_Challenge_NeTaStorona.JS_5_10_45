import type { ILegoRequest, ILegoResonse } from "@/stores/models/LegoDto";
import axios from "axios";
import { LlmSendRequestService } from "@/utils/ServiceAddresses";

export const sendLlmRequest = async (
  request: ILegoRequest,
): Promise<string> => {
  console.log(request);
  try {
    const response = await axios.post<any>(LlmSendRequestService, request);
    return response.data.response;
  } catch (e) {
    throw e;
  }
};
