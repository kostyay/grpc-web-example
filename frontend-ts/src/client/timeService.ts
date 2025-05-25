import { createClient } from "@connectrpc/connect";
import { TimeService } from "./api/time/v1/time_service_pb";
import { createGrpcWebTransport } from "@connectrpc/connect-web";

const timeServiceHost = "http://localhost:8080";

export const getTime = async (): Promise<string> => {
    // Create a gRPC-Web transport and client for the TimeService
    const transport = createGrpcWebTransport({
        baseUrl: timeServiceHost,
    });
    const client = createClient(TimeService, transport)

    // Call the GetCurrentTime method and return the current time
    const response = await client.getCurrentTime({});
    return response.currentTime;
}
