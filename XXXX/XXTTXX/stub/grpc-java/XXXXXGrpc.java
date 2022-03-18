package com.skemaloop.test;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.25.0)",
    comments = "Source: XXXXX.proto")
public final class XXXXXGrpc {

  private XXXXXGrpc() {}

  public static final String SERVICE_NAME = "XXXX.XXTTXX.XXXXX";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.skemaloop.test.XXXXXOuterClass.HelloRequest,
      com.skemaloop.test.XXXXXOuterClass.HelloReply> getSayHelloMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SayHello",
      requestType = com.skemaloop.test.XXXXXOuterClass.HelloRequest.class,
      responseType = com.skemaloop.test.XXXXXOuterClass.HelloReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.skemaloop.test.XXXXXOuterClass.HelloRequest,
      com.skemaloop.test.XXXXXOuterClass.HelloReply> getSayHelloMethod() {
    io.grpc.MethodDescriptor<com.skemaloop.test.XXXXXOuterClass.HelloRequest, com.skemaloop.test.XXXXXOuterClass.HelloReply> getSayHelloMethod;
    if ((getSayHelloMethod = XXXXXGrpc.getSayHelloMethod) == null) {
      synchronized (XXXXXGrpc.class) {
        if ((getSayHelloMethod = XXXXXGrpc.getSayHelloMethod) == null) {
          XXXXXGrpc.getSayHelloMethod = getSayHelloMethod =
              io.grpc.MethodDescriptor.<com.skemaloop.test.XXXXXOuterClass.HelloRequest, com.skemaloop.test.XXXXXOuterClass.HelloReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SayHello"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.skemaloop.test.XXXXXOuterClass.HelloRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.skemaloop.test.XXXXXOuterClass.HelloReply.getDefaultInstance()))
              .setSchemaDescriptor(new XXXXXMethodDescriptorSupplier("SayHello"))
              .build();
        }
      }
    }
    return getSayHelloMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static XXXXXStub newStub(io.grpc.Channel channel) {
    return new XXXXXStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static XXXXXBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new XXXXXBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static XXXXXFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new XXXXXFutureStub(channel);
  }

  /**
   */
  public static abstract class XXXXXImplBase implements io.grpc.BindableService {

    /**
     */
    public void sayHello(com.skemaloop.test.XXXXXOuterClass.HelloRequest request,
        io.grpc.stub.StreamObserver<com.skemaloop.test.XXXXXOuterClass.HelloReply> responseObserver) {
      asyncUnimplementedUnaryCall(getSayHelloMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSayHelloMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.skemaloop.test.XXXXXOuterClass.HelloRequest,
                com.skemaloop.test.XXXXXOuterClass.HelloReply>(
                  this, METHODID_SAY_HELLO)))
          .build();
    }
  }

  /**
   */
  public static final class XXXXXStub extends io.grpc.stub.AbstractStub<XXXXXStub> {
    private XXXXXStub(io.grpc.Channel channel) {
      super(channel);
    }

    private XXXXXStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected XXXXXStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new XXXXXStub(channel, callOptions);
    }

    /**
     */
    public void sayHello(com.skemaloop.test.XXXXXOuterClass.HelloRequest request,
        io.grpc.stub.StreamObserver<com.skemaloop.test.XXXXXOuterClass.HelloReply> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSayHelloMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class XXXXXBlockingStub extends io.grpc.stub.AbstractStub<XXXXXBlockingStub> {
    private XXXXXBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private XXXXXBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected XXXXXBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new XXXXXBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.skemaloop.test.XXXXXOuterClass.HelloReply sayHello(com.skemaloop.test.XXXXXOuterClass.HelloRequest request) {
      return blockingUnaryCall(
          getChannel(), getSayHelloMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class XXXXXFutureStub extends io.grpc.stub.AbstractStub<XXXXXFutureStub> {
    private XXXXXFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private XXXXXFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected XXXXXFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new XXXXXFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.skemaloop.test.XXXXXOuterClass.HelloReply> sayHello(
        com.skemaloop.test.XXXXXOuterClass.HelloRequest request) {
      return futureUnaryCall(
          getChannel().newCall(getSayHelloMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SAY_HELLO = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final XXXXXImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(XXXXXImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SAY_HELLO:
          serviceImpl.sayHello((com.skemaloop.test.XXXXXOuterClass.HelloRequest) request,
              (io.grpc.stub.StreamObserver<com.skemaloop.test.XXXXXOuterClass.HelloReply>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class XXXXXBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    XXXXXBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.skemaloop.test.XXXXXOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("XXXXX");
    }
  }

  private static final class XXXXXFileDescriptorSupplier
      extends XXXXXBaseDescriptorSupplier {
    XXXXXFileDescriptorSupplier() {}
  }

  private static final class XXXXXMethodDescriptorSupplier
      extends XXXXXBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    XXXXXMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (XXXXXGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new XXXXXFileDescriptorSupplier())
              .addMethod(getSayHelloMethod())
              .build();
        }
      }
    }
    return result;
  }
}
