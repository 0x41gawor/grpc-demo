# Intro

## What is: gRPC - 1st approx

gRPC is an open source *remote procedure call* framework made by Google. 

## What is: remote procedure call

https://en.wikipedia.org/wiki/Remote_procedure_call

In [distributed computing](https://en.wikipedia.org/wiki/Distributed_computing), a **remote procedure call** (**RPC**) is when a computer program causes a procedure ([subroutine](https://en.wikipedia.org/wiki/Subroutine)) to execute in a different [address space](https://en.wikipedia.org/wiki/Address_space) (commonly on another computer on a shared network), which is written as if it were a normal (local) procedure call, without the programmer explicitly writing the details for the remote interaction. That is, the programmer writes essentially the same code whether the subroutine is local to the executing program, or remote. This is a form of client–server interaction (caller is client, executor is server), typically implemented via a request–response message-passing system.

The RPC model implies a level of location transparency, namely that calling procedures are largely the same whether they are local or remote, but usually, they are not identical, so local calls can be distinguished from remote calls

RPC is a request–response protocol. An RPC is initiated by the *client*, which sends a request message to a known remote *server* to execute a specified procedure with supplied parameters. 

## What is: gRPC - 2nd approx

**gRPC** (**gRPC Remote Procedure Calls**) is a [cross-platform](https://en.wikipedia.org/wiki/Cross-platform) open source high performance [remote procedure call](https://en.wikipedia.org/wiki/Remote_procedure_call) (RPC) framework. 

It uses [HTTP/2](https://en.wikipedia.org/wiki/HTTP/2) for transport, [Protocol Buffers](https://en.wikipedia.org/wiki/Protocol_Buffers) as the [interface description language](https://en.wikipedia.org/wiki/Interface_description_language), 

and provides features such as:

- authentication, 
- bidirectional streaming and [flow control](https://en.wikipedia.org/wiki/Flow_control_(data)), 
- blocking or nonblocking bindings, 
- cancellation and 
- timeouts.

## What is: Interface Description Language

An **interface description language** or **interface definition language** (**IDL**) is a generic term for a language that lets a program or object written in one language communicate with another program written in an unknown language. IDLs are usually used to describe [data types](https://en.wikipedia.org/wiki/Data_type) and interfaces in a [language-independent](https://en.wikipedia.org/wiki/Language-independent_specification) way, for example, between those written in [C++](https://en.wikipedia.org/wiki/C%2B%2B) and those written in [Java](https://en.wikipedia.org/wiki/Java_(programming_language)).

IDLs are commonly used in [remote procedure call](https://en.wikipedia.org/wiki/Remote_procedure_call) software. In these cases the machines at either end of the *link* may be using different [operating systems](https://en.wikipedia.org/wiki/Operating_system) and computer languages. IDLs offer a bridge between the two different systems.

## What is: Protobuf

**Protocol Buffers** (**Protobuf**) is a [free and open-source](https://en.wikipedia.org/wiki/Free_and_open-source_software) [cross-platform](https://en.wikipedia.org/wiki/Cross-platform_software) data format used to [serialize](https://en.wikipedia.org/wiki/Serialization) structured data.

The method involves:

- an [interface description language](https://en.wikipedia.org/wiki/Interface_description_language) that describes the structure of some data and
- a program that generates source code from that description for generating or parsing a stream of bytes that represents the structured data.

Protocol buffers are Google’s language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

```protobuf
message Person {
  optional string name = 1;
  optional int32 id = 2;
  optional string email = 3;
}
```

It’s like JSON, except it’s smaller and faster, and it generates native language bindings. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

Protocol buffers are a combination of the definition language (created in `.proto` files), the code that the proto compiler generates to interface with data, language-specific runtime libraries, and the serialization format for data that is written to a file (or sent across a network connection).