import React from "react";
import { grpcClient } from "../gprc";
import { useQuery } from "@tanstack/react-query";

export const Board: React.FC = () => {
  const { data } = useQuery({
    queryKey: [""],
    queryFn: () => grpcClient.getGame({ gameId: "abc" }),
  });

  console.log(data);

  return <div>hej</div>;
};
