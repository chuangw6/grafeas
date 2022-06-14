package intoto

import (
	"fmt"
	"time"

	pb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
	intoto "github.com/in-toto/in-toto-golang/in_toto"
	slsa "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v0.2"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToProto function converts the intoto statement with SLSA 0.2 predicate to the
// grafeas proto InTotoStatement where slsa 0.2 predicate is oneof the predicate types
func ToProto(in *intoto.ProvenanceStatement) (*pb.InTotoStatement, error) {
	if in == nil {
		return nil, nil
	}

	subject := convertSubject(in.Subject)
	predicate, err := convertPredicate(in)
	if err != nil {
		return nil, err
	}

	return &pb.InTotoStatement{
		Type:          in.Type,
		Subject:       subject,
		PredicateType: in.PredicateType,
		Predicate:     predicate,
	}, nil

}

func convertPredicate(in *intoto.ProvenanceStatement) (*pb.InTotoStatement_SlsaProvenanceZeroTwo, error) {
	// convert the fields of interface{} type to *structpb.Struct
	parameter, err := toProtoStruct(in.Predicate.Invocation.Parameters)
	if err != nil {
		return nil, err
	}

	buildConfig, err := toProtoStruct(in.Predicate.BuildConfig)
	if err != nil {
		return nil, err
	}

	environment, err := toProtoStruct(in.Predicate.Invocation.Environment)
	if err != nil {
		return nil, err
	}

	return &pb.InTotoStatement_SlsaProvenanceZeroTwo{
		SlsaProvenanceZeroTwo: &pb.SlsaProvenanceZeroTwo{
			Builder:   &pb.SlsaProvenanceZeroTwo_SlsaBuilder{Id: in.Predicate.Builder.ID},
			BuildType: in.Predicate.BuildType,
			Invocation: &pb.SlsaProvenanceZeroTwo_SlsaInvocation{
				ConfigSource: &pb.SlsaProvenanceZeroTwo_SlsaConfigSource{
					Uri:        in.Predicate.Invocation.ConfigSource.URI,
					Digest:     in.Predicate.Invocation.ConfigSource.Digest,
					EntryPoint: in.Predicate.Invocation.ConfigSource.EntryPoint,
				},
				Parameters:  parameter,
				Environment: environment,
			},
			BuildConfig: buildConfig,
			Metadata:    convertMetaData(in.Predicate.Metadata),
			Materials:   convertMaterials(in.Predicate.Materials),
		},
	}, nil
}

func convertSubject(input []intoto.Subject) []*pb.Subject {
	results := []*pb.Subject{}
	for _, subject := range input {
		results = append(results, &pb.Subject{
			Name:   subject.Name,
			Digest: subject.Digest,
		})
	}
	return results
}

func toProtoStruct(input interface{}) (*structpb.Struct, error) {
	if input == nil {
		return nil, nil
	}

	data, err := structpb.NewStruct(input.(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("Unable to convert the interface{} to *structpb.Struct: %v", err)
	}

	return data, nil
}

func convertMetaData(in *slsa.ProvenanceMetadata) *pb.SlsaProvenanceZeroTwo_SlsaMetadata {
	if in == nil {
		return nil
	}

	return &pb.SlsaProvenanceZeroTwo_SlsaMetadata{
		BuildInvocationId: in.BuildInvocationID,
		BuildStartedOn:    convertTime(in.BuildStartedOn),
		BuildFinishedOn:   convertTime(in.BuildFinishedOn),
		Completeness: &pb.SlsaProvenanceZeroTwo_SlsaCompleteness{
			Parameters:  in.Completeness.Parameters,
			Environment: in.Completeness.Environment,
			Materials:   in.Completeness.Materials,
		},
		Reproducible: in.Reproducible,
	}
}

func convertMaterials(in []slsa.ProvenanceMaterial) []*pb.SlsaProvenanceZeroTwo_SlsaMaterial {
	results := []*pb.SlsaProvenanceZeroTwo_SlsaMaterial{}

	for _, material := range in {
		results = append(results, &pb.SlsaProvenanceZeroTwo_SlsaMaterial{
			Uri:    material.URI,
			Digest: material.Digest,
		})
	}
	return results
}

func convertTime(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}
