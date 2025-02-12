package p2p

//Peer is an interface represents remote node 
type Peer interface{

}

/*Transport is anything that can handle the communication 
between the node in the network.This can be from (TCP,UDP,WEBSOCKET ...)
*/
type Transport interface{
  ListenAndAccept() error
}